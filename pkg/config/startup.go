package config

import (
	"github.com/gin-gonic/gin"
)

var motionEngine *gin.Engine

func StartupEnginer(function func()) {
	motionEngine = gin.Default()
	function()
	motionEngine.Run(":8080")
}

func GetMotionEnginer() *gin.Engine {
	return motionEngine
}
