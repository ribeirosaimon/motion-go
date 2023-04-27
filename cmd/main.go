package main

import (
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/repository"
	"github.com/ribeirosaimon/motion-go/version"
)

func main() {
	motionGo := config.NewMotionGo()
	setUpRoles()
	motionGo.AddRouter(version.V1, version.V2)
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
