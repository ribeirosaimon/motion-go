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

func (l *loginController) SignUp(ctx *gin.Context) {
	var signupDto dto.SignUpDto

	if err := ctx.BindJSON(&signupDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	profile, err := l.service.SignUp(ctx, signupDto)
	if err != nil {
		err.Throw(ctx)
		return
	}
	httpResponse.Entity(ctx, http.StatusCreated, profile)
}

func (l *loginController) Login(ctx *gin.Context) {
	var signupDto dto.LoginDto

	if err := ctx.BindJSON(&signupDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	session, err := l.service.Login(ctx, signupDto)
	if err != nil {
		exceptions.InternalServer(err.Error()).Throw(ctx)
		return
	}
	httpResponse.Entity(ctx, http.StatusOK, session)
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
