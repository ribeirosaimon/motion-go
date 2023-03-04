package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/pkg/user"
)

func MotionRouters(engine *gin.Engine) {
	health.NewHeathController(engine)
	user.NewUserController(engine)
}
