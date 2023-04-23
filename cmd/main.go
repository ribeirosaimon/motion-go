package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/internal/config"
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

	health.NewHealthRouter(apiVersion, config.Connect)
	login.NewLoginRouter(apiVersion, config.Connect)
	shoppingcart.NewShoppingCartRouter(apiVersion, config.Connect)
	product.NewProductRouter(apiVersion, config.Connect)
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
