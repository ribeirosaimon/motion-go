package health

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

type healthController struct {
	service *healthService
}

func NewHealthController(service *healthService) healthController {
	return healthController{service: service}
}
func (c healthController) openHealth(ctx *gin.Context) {
	health := c.service.getOpenHealthService()
	httpresponse.Ok(ctx, health)
}

func (c healthController) closeHealth(ctx *gin.Context) {
	health := c.service.getHealthService(middleware.GetLoggedUser(ctx))
	httpresponse.Ok(ctx, health)
}
