package controllers

import (
	"github.com/gin-gonic/gin"
)

type MotionRouter struct {
	Path    string
	Method  string
	Service func(ctx *gin.Context)
}

type motionController struct {
	Engine   *gin.Engine
	Handlers []MotionRouter
}

func NewMotionController(engine *gin.Engine, controllers ...MotionRouter) motionController {
	return motionController{
		Engine:   engine,
		Handlers: controllers,
	}

}
func NewMotionRouter(method, path string, service func(engine *gin.Context)) MotionRouter {
	return MotionRouter{
		Method:  method,
		Path:    path,
		Service: service,
	}
}

func (e motionController) Add() {
	for _, controller := range e.Handlers {
		e.Engine.Handle(controller.Method, controller.Path, controller.Service)
	}
}
