package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/health"
)

func MotionRouters(engine *gin.Engine) {
	health.NewHeathController(engine)
}
