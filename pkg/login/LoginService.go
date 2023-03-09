package login

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/database"
	"github.com/ribeirosaimon/motion-go/pkg/config/http"
	"github.com/ribeirosaimon/motion-go/pkg/exceptions"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"github.com/ribeirosaimon/motion-go/repository"
)

type loginService struct {
	userRepository repository.MotionRepository[domain.MotionUser]
}

func NewLoginService() loginService {
	userRepository := repository.NewUserRepository(database.Connect())
	return loginService{userRepository: userRepository}
}

func (l loginService) loginUserService(c *gin.Context) {
	var body loginDto

	if err := c.Bind(&body); err != nil {
		exceptions.BodyError(c)
	}
	id, err := l.userRepository.FindById(8)
	if err != nil {
		return
	}
	err = security.CheckPassword(body.Password, id.Password)
	if err != nil {
		fmt.Println(err)
	}
	http.Created(c, "deu bom")
}

func (l loginService) signUpService(c *gin.Context) {
	var body signUpDto
	if err := c.Bind(&body); err != nil {
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
	http.Created(c, savedUser)
}

type signUpDto struct {
	loginDto
	Name string `json:"name"`
}

type loginDto struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
