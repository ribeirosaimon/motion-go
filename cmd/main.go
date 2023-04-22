package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config"
	"github.com/ribeirosaimon/motion-go/pkg/routers"
	"github.com/ribeirosaimon/motion-go/repository"
)

var motionEngine *gin.Engine

func main() {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	setUpRoles()
	motionEngine = gin.Default()
	routers.MotionRouters(motionEngine)
	serverPort := p.GetInt("server.port", 8080)
	motionEngine.Run(fmt.Sprintf(":%d", serverPort))
}

func setUpRoles() {
	connect, close := config.Connect()
	roleRepository := repository.NewRoleRepository(connect)
	allRoles := []domain.RoleEnum{domain.USER, domain.ADMIN}
	for _, i := range allRoles {
		existByName := roleRepository.ExistByField("name", i)
		if !existByName {
			roleRepository.Save(domain.Role{Name: i})
		}

	}
	close.Close()
}
