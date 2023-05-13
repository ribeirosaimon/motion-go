package main

import (
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/health"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/login"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/product"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

func main() {
	motionGo := config.NewMotionGo()

	db.Conn.InitializeDatabases()

	setUpRoles()
	motionGo.AddRouter(version1, version2)
	motionGo.RunEngine(motionGo.PropertiesFile.GetInt("server.port.baseapp", 0))
}

func setUpRoles() {
	connect := db.Conn.GetPostgreSQL()
	defer db.Conn.ClosePostgreSQL()

	roleRepository := repository.NewRoleRepository(connect)
	allRoles := []sqlDomain.RoleEnum{sqlDomain.USER, sqlDomain.ADMIN}
	for _, i := range allRoles {
		existByName := roleRepository.ExistByField("name", i)
		if !existByName {
			roleRepository.Save(sqlDomain.Role{Name: i})
		}

	}

}

var version1 = config.RoutersVersion{
	Version: "v1",
	Handlers: []func(conn *db.Connections) config.MotionController{
		//health.NewHealthRouter,
		login.NewLoginRouter,
		product.NewProductRouter,
	},
}

var version2 = config.RoutersVersion{
	Version: "v2",
	Handlers: []func(conn *db.Connections) config.MotionController{
		health.NewHealthRouter,
	},
}
