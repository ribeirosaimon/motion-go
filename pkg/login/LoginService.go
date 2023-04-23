package login

import (
	"database/sql"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"gorm.io/gorm"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/profile"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"github.com/ribeirosaimon/motion-go/pkg/session"
	"github.com/ribeirosaimon/motion-go/repository"
)

type loginService struct {
	userRepository repository.MotionRepository[domain.MotionUser]
	profileService profile.Service
	sessionService session.Service
	closeDb        *sql.DB
}

func NewLoginService(conn *gorm.DB, close *sql.DB) loginService {
	return loginService{
		userRepository: repository.NewUserRepository(conn),
		profileService: profile.NewProfileService(conn, close),
		sessionService: session.NewLoginService(conn, close),
		closeDb:        close,
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

func (l loginService) signUpService(signupDto SignUpDto) (domain.Profile, *exceptions.Error) {

	if signupDto.Email == "" {
		return domain.Profile{}, exceptions.FieldError("email")
	}

	if signupDto.Email == "" {
		return domain.Profile{}, exceptions.FieldError("password")
	}

	var user domain.MotionUser
	password, err := security.EncryptPassword(signupDto.Password)
	if err != nil {
		return domain.Profile{}, exceptions.BodyError()
	}
	user.Name = signupDto.Name
	user.Password = password
	user.Email = signupDto.Email

	savedUser, err := l.userRepository.Save(user)
	if err != nil {
		return domain.Profile{}, exceptions.InternalServer(err.Error())
	}

	profileUser, err := l.profileService.SaveProfileUser(savedUser, signupDto.Roles)
	if err != nil {
		return domain.Profile{}, exceptions.InternalServer(err.Error())
	}
	return profileUser, nil
}
