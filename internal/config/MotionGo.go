package config

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type motionGo struct {
	MotionEngine   *gin.Engine
	SqlDatabase    func() (*gorm.DB, *sql.DB)
	MongoDatabase  *mongo.Client
	PropertiesFile *properties.Properties
	Routers        []MotionController
}

func NewMotionGo() motionGo {
	return motionGo{
		MotionEngine:   gin.Default(),
		SqlDatabase:    db.ConnectSqlDb,
		MongoDatabase:  db.ConnectNoSqlDb(),
		PropertiesFile: properties.MustLoadFile("config.properties", properties.UTF8),
	}
}

func (m *motionGo) AddRouter(routers ...MotionController) {
	//apiVersion := m.MotionEngine.Group(fmt.Sprintf("/api/%s", version))

	m.Routers = append(m.Routers, routers...)
}

func (m *motionGo) RunEngine(serverPort int) {
	for _, routers := range m.Routers {
		for _, controller := range routers.Handlers {
			handlerFunc := gin.HandlerFunc(controller.Service)
			controller.Middleware = append(controller.Middleware, handlerFunc)
			m.MotionEngine.Handle(controller.Method, controller.Path, controller.Middleware...)
		}
	}
	fmt.Sprintf(":%d", serverPort)
}
