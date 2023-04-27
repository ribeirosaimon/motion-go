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
func NewMotionRouter(method, path string, service func(*gin.Context),
	middleware ...gin.HandlerFunc) MotionRouter {
	return MotionRouter{
		Method:     method,
		Path:       path,
		Service:    service,
		Middleware: middleware,
	}
}

//
// func (e motionController) Add() {
//
//	for _, controller := range e.Handlers {
//		handlerFunc := gin.HandlerFunc(controller.Service)
//		controller.Middleware = append(controller.Middleware, handlerFunc)
//		e.Engine.Handle(controller.Method, controller.Path, controller.Middleware...)
//	}
// }
