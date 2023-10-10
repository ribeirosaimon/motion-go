package service

import (
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/src/pkg/dto"
)

type AuthService struct {
	userRepository *repository.MotionSQLRepository[sqlDomain.MotionUser]
	profileService *ProfileService
	sessionService *SessionService
}

func NewAuthService(conn *db.Connections) AuthService {
	return AuthService{
		userRepository: repository.NewUserRepository(conn.GetPgsqTemplate()),
		profileService: NewProfileService(conn),
		sessionService: NewSessionService(conn),
	}
}

func (l *AuthService) Login(loginDto dto.LoginDto) (string, *exceptions.Error) {
	user, err := l.userRepository.FindByField("email", loginDto.Email)
	if err != nil {
		return "", exceptions.NotFound()
	}
	err = middleware.CheckPassword(loginDto.Password, user.Password)
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

func (l *AuthService) SignUp(signupDto dto.SignUpDto) (sqlDomain.Profile, *exceptions.Error) {

	if signupDto.Email == "" {
		return sqlDomain.Profile{}, exceptions.FieldError("email")
	}

	if signupDto.Email == "" {
		return sqlDomain.Profile{}, exceptions.FieldError("password")
	}

	var user sqlDomain.MotionUser
	password, err := middleware.EncryptPassword(signupDto.Password)
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

func (l *AuthService) WhoAmI(userId uint64) (sqlDomain.Profile, error) {

	user, err := l.profileService.FindProfileByUserId(userId)
	if err != nil {
		return sqlDomain.Profile{}, err
	}
	return user, nil
}
