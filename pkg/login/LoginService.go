package login

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/database"
	"github.com/ribeirosaimon/motion-go/pkg/config/http"
	"github.com/ribeirosaimon/motion-go/pkg/exceptions"
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

func NewLoginService() loginService {
	connect, s := database.Connect()
	return loginService{
		userRepository: repository.NewUserRepository(connect),
		profileService: profile.NewProfileService(),
		sessionService: session.NewLoginService(),
		closeDb:        s,
	}
}

func (l loginService) loginUserService(c *gin.Context) {
	var body LoginDto

	if err := c.BindJSON(&body); err != nil {
		exceptions.BodyError(c)
	}
	user, err := l.userRepository.FindByField("email", body.Email)
	if err != nil {
		exceptions.Unauthorized(c)
		return
	}
	err = security.CheckPassword(body.Password, user.Password)
	if err != nil {
		exceptions.InternalServer(c, err.Error())
		return
	}
	user.LoginCount += 1
	user.LastLogin = time.Now()
	savedUser, err := l.userRepository.Save(user)
	if err != nil {
		exceptions.InternalServer(c, err.Error())
		return
	}
	profileUser, err := l.profileService.FindProfileByUserId(savedUser.Id)
	if err != nil {
		exceptions.InternalServer(c, err.Error())
		return
	}

	userSession, err := l.sessionService.SaveUserSession(profileUser)
	if err != nil {
		exceptions.InternalServer(c, err.Error())
		return
	}

	http.Created(c, userSession.SessionId)
}

func (l loginService) signUpService(c *gin.Context) {
	var body SignUpDto
	defer l.closeDb.Close()

	if err := c.BindJSON(&body); err != nil {
		exceptions.BodyError(c)
	}
	if body.Email == "" {
		exceptions.FieldError(c, "email")
		return
	}
	if body.Password == "" {
		exceptions.FieldError(c, "password")
		return
	}
	var user domain.MotionUser
	password, err := security.EncryptPassword(body.Password)
	if err != nil {
		exceptions.FieldError(c, "password")
		return
	}
	user.Name = body.Name
	user.Password = password
	user.Email = body.Email

	savedUser, err := l.userRepository.Save(user)
	if err != nil {
		exceptions.BodyError(c)
		return
	}

	profileUser, err := l.profileService.SaveProfileUser(savedUser)
	if err != nil {
		exceptions.BodyError(c)
		return
	}
	http.Created(c, profileUser)
}

type SignUpDto struct {
	LoginDto
	Name string `json:"name"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
