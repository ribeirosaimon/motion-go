package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/internal/db"
)

type RoutersVersion struct {
	Version  string
	Handlers []func(*db.Connections) MotionController
}

type motionGo struct {
	MotionEngine   *gin.Engine
	Connections    *db.Connections
	PropertiesFile *properties.Properties
	Routers        []RoutersVersion
}

func NewMotionGo() motionGo {
	return motionGo{
		MotionEngine:   gin.Default(),
		PropertiesFile: properties.MustLoadFile("config.properties", properties.UTF8),
	}
}

func (m *motionGo) AddRouter(version ...RoutersVersion) {
	m.Routers = append(m.Routers, version...)
}

func (m *motionGo) RunEngine(serverPort int) {
	for _, routerVersions := range m.Routers {
		apiVersion := m.MotionEngine.Group(fmt.Sprintf("/api/%s", routerVersions.Version))
		for _, routersFunc := range routerVersions.Handlers {
			routers := routersFunc(m.Connections)
			pathEngineer := apiVersion.Group(routers.Path)
			for _, controller := range routers.Handlers {
				handlerFunc := gin.HandlerFunc(controller.Service)
				controller.Middleware = append(controller.Middleware, handlerFunc)
				pathEngineer.Handle(controller.Method, controller.Path, controller.Middleware...)
			}
		}
	}

	m.MotionEngine.Run(fmt.Sprintf(":%d", serverPort))
}
