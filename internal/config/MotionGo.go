package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
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
	gin.DefaultWriter = ioutil.Discard
	return &MotionGo{
		MotionEngine:   gin.New(),
		PropertiesFile: properties.MustLoadFile(propertiesFile, properties.UTF8),
	}
}

func (m *MotionGo) AddRouter(version ...RoutersVersion) {
	m.Routers = append(m.Routers, version...)
}

func (m *MotionGo) CreateRouters() {
	m.MotionEngine.Use(middleware.NewLogger("motion"))
	for _, routerVersions := range m.Routers {
		apiVersion := m.MotionEngine.Group(fmt.Sprintf("/api/%s", routerVersions.Version))
		for _, routersFunc := range routerVersions.Handlers {
			routers := routersFunc()
			pathEngineer := apiVersion.Group(routers.Path)

			for _, controller := range routers.Handlers {
				if m.MotionEngine.Routes() != nil {
					for _, routersAdded := range m.MotionEngine.Routes() {
						if !(routersAdded.Path == routers.Path && routersAdded.Method == controller.Method) {
							addHandlerToEngine(controller, pathEngineer)
						}
					}
				} else {
					addHandlerToEngine(controller, pathEngineer)
				}

			}
		}
	}

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
