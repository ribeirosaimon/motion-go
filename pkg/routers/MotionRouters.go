package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/database"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/pkg/login"
)

func MotionRouters(engine *gin.Engine) {
	health.NewHealthRouter(engine, database.Connect)
	login.NewLoginRouter(engine, database.Connect)
}
