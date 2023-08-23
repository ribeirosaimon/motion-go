package config

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
)

type RoutersVersion struct {
	Version  string
	Handlers []func() MotionController
}

type MotionGo struct {
	MotionEngine   *gin.Engine
	PropertiesFile *properties.Properties
	Routers        []RoutersVersion
}

func NewMotionGo(propertiesFile string) *MotionGo {
	// gin.DefaultWriter = ioutil.Discard
	var engine *gin.Engine
	engine = gin.New()
	return &MotionGo{
		MotionEngine:   engine,
		PropertiesFile: properties.MustLoadFile(propertiesFile, properties.UTF8),
	}
}

func (m *MotionGo) AddRouter(version ...RoutersVersion) {
	m.Routers = append(m.Routers, version...)
}

func (m *MotionGo) CreateRouters(logger func() gin.HandlerFunc) {
	m.MotionEngine.Use(logger())
	for _, routerVersions := range m.Routers {
		apiVersion := m.MotionEngine.Group(fmt.Sprintf("/api/%s", routerVersions.Version))
		for _, routersFunc := range routerVersions.Handlers {
			routers := routersFunc()
			pathEngineer := apiVersion.Group(routers.Path)

			for _, controller := range routers.Handlers {
				if m.MotionEngine.Routes() != nil {
					path := fmt.Sprintf("%s%s", pathEngineer.BasePath(), controller.Path)

					if !existRouter(m.MotionEngine.Routes(), path, controller.Method) {
						addHandlerToEngine(controller, pathEngineer)
					}
				} else {
					addHandlerToEngine(controller, pathEngineer)
				}

			}
		}
	}

}

func existRouter(routes gin.RoutesInfo, path string, method string) bool {
	for _, v := range routes {
		if v.Path == path && v.Method == method {
			return true
		}
	}
	return false
}

func addHandlerToEngine(controller MotionRouter, pathEngineer *gin.RouterGroup) {
	handlerFunc := gin.HandlerFunc(controller.Service)

	controller.Middleware = append(controller.Middleware, handlerFunc)
	pathEngineer.Handle(controller.Method, controller.Path, controller.Middleware...)
	log.Printf("Add %s with path: %s", controller.Method, controller.Path)
}

func (m *MotionGo) RunEngine(serverPort int) {
	m.MotionEngine.Run(fmt.Sprintf(":%d", serverPort))
}
