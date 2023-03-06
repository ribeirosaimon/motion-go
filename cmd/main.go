package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/routers"
)

var motionEngine *gin.Engine

func main() {
	motionEngine = gin.Default()

	routers.MotionRouters(motionEngine)
	motionEngine.Run(":8080")
}
