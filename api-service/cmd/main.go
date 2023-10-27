package main

import (
	"fmt"
	"github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/grpcconnection"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/internal/router"
	"github.com/ribeirosaimon/motion-go/internal/util"
)

func main() {
	propertiesFile := "config.properties"
	dir, _ := util.FindRootDir()

	motionGo := config.NewMotionGo(fmt.Sprintf("%s/%s", dir, propertiesFile))

	motionGo.MotionEngine.Use(middleware.CorsMiddleware)
	db.Conn = &db.Connections{}
	db.Conn.InitializeDatabases(motionGo.PropertiesFile)
	grpcconnection.NewConnection(motionGo.PropertiesFile.GetString("grpc.host", ""),
		motionGo.PropertiesFile.GetString("grpc.port", ""))

	config.NewMotionConfig(motionGo.PropertiesFile)
	setUpRoles()
	go middleware.NewMotionCache(db.Conn)

	motionGo.AddRouter(version1)
	motionGo.CreateRouters(middleware.NewLogger)
	motionGo.RunEngine(motionGo.PropertiesFile.GetInt("server.port.src", 0))
}

func setUpRoles() {
	connect := db.Conn.GetPgsqTemplate()
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
	Handlers: []func() config.MotionController{
		router.NewHealthRouter,
		router.NewLoginRouter,
		router.NewCompanyRouter,
		router.NewPortfolioRouter,
		router.NewTransactionRouter,
	},
}

var version2 = config.RoutersVersion{
	Version: "v2",
	Handlers: []func() config.MotionController{
		router.NewHealthRouter,
		router.NewHealthRouter,
	},
}
