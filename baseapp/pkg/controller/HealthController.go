package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/service"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

type healthController struct {
	service *service.HealthService
}

func NewHealthController() *healthController {
	healthService := service.NewHealthService()

	return &healthController{service: &healthService}
}
func (c *healthController) OpenHealth(ctx *gin.Context) {
	health := c.service.GetOpenHealthService()

	httpresponse.Ok(ctx, health)
}

func (c *healthController) CloseHealth(ctx *gin.Context) {
	user, err := middleware.GetLoggedUser(ctx)
	if err != nil {
		exceptions.Forbidden().Throw(ctx)
		return
	}
	health := c.service.GetHealthService(user)
	httpresponse.Ok(ctx, health)
}
