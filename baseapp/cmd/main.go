package main

import (
	"context"
	"fmt"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/router"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/internal/util"
)

func main() {
	propertiesFile := "config.properties"
	dir, _ := util.FindRootDir()

	motionGo := config.NewMotionGo(fmt.Sprintf("%s/%s", dir, propertiesFile))

	db.Conn = &db.Connections{}
	db.Conn.InitializeDatabases(motionGo.PropertiesFile)

	motionConfig := config.NewMotionConfig(context.Background(), motionGo.PropertiesFile)
	setUpRoles()

	go middleware.NewMotionCache(db.Conn, motionConfig.HaveScraping, motionConfig.ScrapingTime, motionConfig.CacheTime)

	motionGo.AddRouter(version1)
	motionGo.CreateRouters()
	motionGo.RunEngine(motionGo.PropertiesFile.GetInt("server.port.baseapp", 0))
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
	},
}

var version2 = config.RoutersVersion{
	Version: "v2",
	Handlers: []func() config.MotionController{
		router.NewHealthRouter,
		router.NewHealthRouter,
	},
}
