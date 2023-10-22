package service

import (
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"

	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/dto"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/internal/util"
	"github.com/ribeirosaimon/motion-go/src/emailSender"
)

type AuthService struct {
	userRepository     *repository.MotionSQLRepository[sqlDomain.MotionUser]
	profileService     *ProfileService
	sessionService     *SessionService
	transactionService *TransactionService
}

func NewAuthService(conn *db.Connections) AuthService {
	return AuthService{
		userRepository:     repository.NewUserRepository(conn.GetPgsqTemplate()),
		profileService:     NewProfileService(conn),
		sessionService:     NewSessionService(conn),
		transactionService: NewTransactionService(conn),
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
	if profileUser.Status == sqlDomain.INACTIVE {
		return "", exceptions.Unauthorized()
	}
	if err != nil {
		return "", exceptions.InternalServer(err.Error())
	}

	userSession, err := l.sessionService.SaveUserSession(profileUser)
	if err != nil {
		return "", exceptions.InternalServer(err.Error())
	}
	return userSession.Id, nil
}

func (l *AuthService) SignUp(signupDto dto.SignUpDto) (sqlDomain.Profile, *exceptions.Error) {

	if signupDto.Email == "" || !util.ValidateEmail(signupDto.Email) {
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

	go emailSender.SendEmail(profileUser.Code)

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

func (l *AuthService) ValidateEmail(loggedUser middleware.LoggedUser, code string) error {
	profile, err := l.WhoAmI(loggedUser.ProfileId)
	if err != nil {
		return err
	}
	if profile.Status == sqlDomain.ACTIVE {
		return errors.New("this profile was active")
	}
	if profile.Code != code {
		return errors.New("this code was wrong")
	}
	profile.Status = sqlDomain.ACTIVE
	profile.UpdatedAt = time.Now()

	l.transactionService.Deposit(loggedUser, dto.Deposit{Value: float64(config.GetMotionConfig().InitialValue)})
	profile, err = l.profileService.profileRepository.Save(profile)
	if err != nil {
		return err
	}
	return nil
}
