package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpResponse"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/src/pkg/service"
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

	httpResponse.Entity(ctx, http.StatusOK, health)
}

func (c *healthController) CloseHealth(ctx *gin.Context) {
	user, err := middleware.GetLoggedUser(ctx)
	if err != nil {
		exceptions.Forbidden().Throw(ctx)
		return
	}
	health := c.service.GetHealthService(user)
	httpResponse.Entity(ctx, http.StatusOK, health)
}

func (c *healthController) GetConfigValue(ctx *gin.Context) {
	config := c.service.GetConfigurations()
	httpResponse.Entity(ctx, http.StatusOK, config)
}
