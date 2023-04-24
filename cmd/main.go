package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/pkg/login"
	"github.com/ribeirosaimon/motion-go/pkg/product"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"
	"github.com/ribeirosaimon/motion-go/repository"
)

var motionEngine *gin.Engine

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

	health.NewHealthRouter(apiVersion, db.ConnectSqlDb)
	login.NewLoginRouter(apiVersion, db.ConnectSqlDb)
	shoppingcart.NewShoppingCartRouter(apiVersion, db.ConnectSqlDb)
	product.NewProductRouter(apiVersion, db.ConnectSqlDb)
}

func setUpRoles() {
	connect, close := db.ConnectSqlDb()
	roleRepository := repository.NewRoleRepository(connect)
	allRoles := []sqlDomain.RoleEnum{sqlDomain.USER, sqlDomain.ADMIN}
	for _, i := range allRoles {
		existByName := roleRepository.ExistByField("name", i)
		if !existByName {
			roleRepository.Save(sqlDomain.Role{Name: i})
		}

	}
	close.Close()
}
