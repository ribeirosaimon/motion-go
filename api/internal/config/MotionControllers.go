package config

import (
	"github.com/gin-gonic/gin"
)

type MotionRouter struct {
	Path       string
	Method     string
	Service    func(*gin.Context)
	Middleware []gin.HandlerFunc
}

type MotionController struct {
	Path     string
	Handlers []MotionRouter
}

func NewMotionController(path string, controllers ...MotionRouter) MotionController {
	return MotionController{
		Path:     path,
		Handlers: controllers,
	}

}
func NewMotionRouter(method, path string,
	service gin.HandlerFunc,
	middleware ...gin.HandlerFunc) MotionRouter {
	return MotionRouter{
		Method:     method,
		Path:       path,
		Service:    service,
		Middleware: middleware,
	}
}
