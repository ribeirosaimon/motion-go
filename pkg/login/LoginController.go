package login

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
)

type loginController struct {
	service *loginService
}

func NewLoginControler(service *loginService) loginController {
	return loginController{service: service}
}

func (l loginController) signUp(c *gin.Context) {
	var signupDto SignUpDto

	if err := c.BindJSON(&signupDto); err != nil {
		exceptions.BodyError().Throw(c)
		return
	}
	profile, err := l.service.signUpService(signupDto)
	if err != nil {
		err.Throw(c)
		return
	}
	httpresponse.Created(c, profile)
}

func (l loginController) login(c *gin.Context) {
	var signupDto LoginDto

	if err := c.BindJSON(&signupDto); err != nil {
		exceptions.BodyError().Throw(c)
		return
	}
	session, err := l.service.loginService(signupDto)
	if err != nil {
		err.Throw(c)
		return
	}
	httpresponse.Created(c, session)
}

type SignUpDto struct {
	LoginDto
	Name  string            `json:"name"`
	Roles []domain.RoleEnum `json:"roles"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
