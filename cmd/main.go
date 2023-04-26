package main

import (
	"github.com/ribeirosaimon/motion-go/pkg/login"
	"github.com/ribeirosaimon/motion-go/pkg/product"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"

	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/repository"
)

// var motionEngine *gin.Engine

func main() {
	motionGo := config.NewMotionGo()
	setUpRoles()
	motionGo.AddRouter(
		//motionGo.PropertiesFile.GetString("api.version", "v1"),
		health.NewHealthRouter(db.ConnectSqlDb),
		login.NewLoginRouter(db.ConnectSqlDb),
		shoppingcart.NewShoppingCartRouter(db.ConnectSqlDb),
		product.NewProductRouter(db.ConnectSqlDb),
	)

	motionGo.RunEngine(motionGo.PropertiesFile.GetInt("server.port", 8080))
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
