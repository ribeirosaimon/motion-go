package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/database"
	"github.com/ribeirosaimon/motion-go/pkg/routers"
	"github.com/ribeirosaimon/motion-go/repository"
)

var motionEngine *gin.Engine

func main() {
	setUpRoles()
	motionEngine = gin.Default()
	routers.MotionRouters(motionEngine)
	motionEngine.Run(":8080")
}

func setUpRoles() {
	connect, s := database.Connect()
	defer s.Close()
	roleRepository := repository.NewRoleRepository(connect)
	allRoles := []domain.RoleEnum{domain.ADMIN, domain.USER}
	for _, i := range allRoles {
		_, err := roleRepository.FindByField("name", i)
		if err != nil {
			roleRepository.Save(domain.Role{Name: i})
		}

	}
}
