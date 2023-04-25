package config

import (
	"database/sql"

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
	Routers        []MotionRouter
}

func NewMotionGo() motionGo {
	return motionGo{
		MotionEngine:   gin.Default(),
		SqlDatabase:    db.ConnectSqlDb,
		MongoDatabase:  db.ConnectNoSqlDb(),
		PropertiesFile: properties.MustLoadFile("config.properties", properties.UTF8),
	}
}

func (m *motionGo) AddRouter(routers ...MotionRouter) {
	m.Routers = append(m.Routers, routers...)
}
