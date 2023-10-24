package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/response"
	"github.com/ribeirosaimon/motion-go/internal/service"
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

	response.Entity(ctx, http.StatusOK, health)
}

func (c *healthController) CloseHealth(ctx *gin.Context) {
	loggedUser := middleware.GetLoggedUser(ctx)
	health := c.service.GetHealthService(loggedUser)
	response.Entity(ctx, http.StatusOK, health)
}

func (c *healthController) GetConfigValue(ctx *gin.Context) {
	config := c.service.GetConfigurations()
	response.Entity(ctx, http.StatusOK, config)
}
