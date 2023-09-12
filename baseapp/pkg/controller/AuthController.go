package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/dto"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/service"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpResponse"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

type loginController struct {
	service *service.AuthService
}

func NewAuthController() *loginController {
	authService := service.NewAuthService(db.Conn)
	return &loginController{service: &authService}
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
	httpResponse.Entity(c, http.StatusCreated, profile)
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
	httpResponse.Entity(c, http.StatusOK, session)
}

func (l *loginController) WhoAmI(c *gin.Context) {
	user, err := middleware.GetLoggedUser(c)
	if err != nil {
		exceptions.Forbidden().Throw(c)
		return
	}
	i, err := l.service.WhoAmI(user.UserId)
	if err != nil {
		exceptions.Forbidden().Throw(c)
		return
	}
	httpResponse.Entity(c, http.StatusOK, i)
}
