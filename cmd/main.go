package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/domain/sql"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/pkg/login"
	"github.com/ribeirosaimon/motion-go/pkg/product"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"
	"github.com/ribeirosaimon/motion-go/repository"
)

var motionEngine *gin.Engine

type MotionGo struct {
	MongoTemplate *mongo.Client
	EntityManager *gorm.DB
	GinEngine     *gin.Engine
}

func main() {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	setUpRoles()
	motionEngine = gin.Default()
	motionRouters(motionEngine)
	serverPort := p.GetInt("server.port", 8080)
	motionEngine.Run(fmt.Sprintf(":%d", serverPort))
}

func motionRouters(engine *gin.Engine) {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	apiVersion := engine.Group(fmt.Sprintf("/api/%s", p.GetString("api.version", "v1")))

	health.NewHealthRouter(apiVersion, config.ConnectSqlDb)
	login.NewLoginRouter(apiVersion, config.ConnectSqlDb)
	shoppingcart.NewShoppingCartRouter(apiVersion, config.ConnectSqlDb)
	product.NewProductRouter(apiVersion, config.ConnectSqlDb)
}

func setUpRoles() {
	connect, close := config.ConnectSqlDb()
	roleRepository := repository.NewRoleRepository(connect)
	allRoles := []sql.RoleEnum{sql.USER, sql.ADMIN}
	for _, i := range allRoles {
		existByName := roleRepository.ExistByField("name", i)
		if !existByName {
			roleRepository.Save(sql.Role{Name: i})
		}

	}
	close.Close()
}
