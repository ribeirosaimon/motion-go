package login

import (
	"time"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/profile"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/session"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/internal/security"
)

type loginService struct {
	userRepository repository.MotionRepository[sqlDomain.MotionUser]
	profileService profile.Service
	sessionService session.Service
}

func NewLoginService(conn *db.Connections) loginService {
	return loginService{
		userRepository: repository.NewUserRepository(conn.sqlStruct.Conn),
		profileService: profile.NewProfileService(conn),
		sessionService: session.NewLoginService(conn),
	}
}

func (l loginService) loginService(loginDto LoginDto) (string, *exceptions.Error) {
	user, err := l.userRepository.FindByField("email", loginDto.Email)
	if err != nil {
		return "", exceptions.NotFound()
	}
	err = security.CheckPassword(loginDto.Password, user.Password)
	if err != nil {
		return "", exceptions.FieldError("password")
	}
	user.LoginCount += 1
	user.LastLogin = time.Now()
	savedUser, err := l.userRepository.Save(user)
	if err != nil {
		return "", exceptions.Unauthorized()
	}
	profileUser, err := l.profileService.FindProfileByUserId(savedUser.Id)
	if err != nil {
		return "", exceptions.InternalServer(err.Error())
	}

	userSession, err := l.sessionService.SaveUserSession(profileUser)
	if err != nil {
		return "", exceptions.InternalServer(err.Error())
	}
	return userSession.SessionId, nil
}

func (l loginService) signUpService(signupDto SignUpDto) (sqlDomain.Profile, *exceptions.Error) {

	if signupDto.Email == "" {
		return sqlDomain.Profile{}, exceptions.FieldError("email")
	}

	if signupDto.Email == "" {
		return sqlDomain.Profile{}, exceptions.FieldError("password")
	}

	var user sqlDomain.MotionUser
	password, err := security.EncryptPassword(signupDto.Password)
	if err != nil {
		return sqlDomain.Profile{}, exceptions.BodyError()
	}
	user.Name = signupDto.Name
	user.Password = password
	user.Email = signupDto.Email

	savedUser, err := l.userRepository.Save(user)
	if err != nil {
		return sqlDomain.Profile{}, exceptions.InternalServer(err.Error())
	}

	profileUser, err := l.profileService.SaveProfileUser(savedUser, signupDto.Roles)
	if err != nil {
		return sqlDomain.Profile{}, exceptions.InternalServer(err.Error())
	}
	return profileUser, nil
}
