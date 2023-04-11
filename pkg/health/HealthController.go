package health

import (
	"github.com/gin-gonic/gin"
	motionHttp "github.com/ribeirosaimon/motion-go/pkg/httpresponse"
	"github.com/ribeirosaimon/motion-go/pkg/security"
)

type healthController struct {
	service *healthService
}

func NewHealthController(service *healthService) healthController {
	return healthController{service: service}
}
func (c healthController) openHealth(ctx *gin.Context) {
	health := c.service.getOpenHealthService()
	motionHttp.Ok(ctx, health)
}

func (c healthController) closeHealth(ctx *gin.Context) {
	health := c.service.getHealthService(security.GetLoggedUser(ctx))
	motionHttp.Ok(ctx, health)
}
