package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/pkg/login"
	"github.com/ribeirosaimon/motion-go/pkg/product"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"
	"github.com/ribeirosaimon/motion-go/repository"
)

// var motionEngine *gin.Engine

func main() {
	motionGo := config.NewMotionGo()
	setUpRoles()
	motionRouters(motionGo.MotionEngine)
	serverPort := motionGo.PropertiesFile.GetInt("server.port", 8080)
	motionGo.MotionEngine.Run(fmt.Sprintf(":%d", serverPort))
}

func motionRouters(engine *gin.Engine) {

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
