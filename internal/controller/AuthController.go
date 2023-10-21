package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/dto"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/response"
	"github.com/ribeirosaimon/motion-go/internal/service"
)

type loginController struct {
	service *service.AuthService
}

func NewAuthController() *loginController {
	authService := service.NewAuthService(db.Conn)
	return &loginController{service: &authService}
}

func (l *loginController) SignUp(ctx *gin.Context) {
	var signupDto dto.SignUpDto

	if err := ctx.BindJSON(&signupDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	profile, err := l.service.SignUp(signupDto)
	if err != nil {
		err.Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusCreated, profile)
}

func (l *loginController) Login(ctx *gin.Context) {
	var signupDto dto.LoginDto

	if err := ctx.BindJSON(&signupDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	session, err := l.service.Login(signupDto)
	if err != nil {
		exceptions.NotFound().Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusOK, session)
}

func (l *loginController) WhoAmI(c *gin.Context) {
	loggedUser := middleware.GetLoggedUser(c)

	i, err := l.service.WhoAmI(loggedUser.UserId)
	if err != nil {
		exceptions.Forbidden().Throw(c)
		return
	}
	response.Entity(c, http.StatusOK, i)
}
func (l *loginController) ValidateEmail(ctx *gin.Context) {
	loggedUser := middleware.GetLoggedUser(ctx)

	var code dto.ValidateEmailDto
	if err := ctx.BindJSON(&code); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	err := l.service.ValidateEmail(loggedUser, code.Code)
	if err != nil {
		exceptions.MotionError("This code was wrong").Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusOK, nil)
}
