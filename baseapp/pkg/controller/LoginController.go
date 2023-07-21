package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/dto"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/service"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
)

type loginController struct {
	service *service.LoginService
}

func NewLoginControler() *loginController {
	loginService := service.NewLoginService(db.Conn)
	return &loginController{service: &loginService}
}

func (l *loginController) SignUp(c *gin.Context) {
	var signupDto dto.SignUpDto

	if err := c.BindJSON(&signupDto); err != nil {
		exceptions.BodyError().Throw(c)
		return
	}
	profile, err := l.service.SignUp(signupDto)
	if err != nil {
		err.Throw(c)
		return
	}
	httpresponse.Created(c, profile)
}

func (l *loginController) Login(c *gin.Context) {
	var signupDto dto.LoginDto

	if err := c.BindJSON(&signupDto); err != nil {
		exceptions.BodyError().Throw(c)
		return
	}
	session, err := l.service.Login(signupDto)
	if err != nil {
		err.Throw(c)
		return
	}
	httpresponse.Ok(c, session)
}
