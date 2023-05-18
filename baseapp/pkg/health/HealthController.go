package health

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

type healthController struct {
	service *healthService
}

func NewHealthController(service *healthService) healthController {
	return healthController{service: service}
}
func (c healthController) OpenHealth(ctx *gin.Context) {
	health := c.service.getOpenHealthService()
	httpresponse.Ok(ctx, health)
}

func (c healthController) CloseHealth(ctx *gin.Context) {
	user, err := middleware.GetLoggedUser(ctx)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	health := c.service.getHealthService(user)
	httpresponse.Ok(ctx, health)
}
