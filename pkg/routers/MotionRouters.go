package routers

import (
	"fmt"
	"github.com/ribeirosaimon/motion-go/pkg/product"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/pkg/config"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/pkg/login"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"
)

func MotionRouters(engine *gin.Engine) {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	apiVersion := engine.Group(fmt.Sprintf("/api/%s", p.GetString("api.version", "v1")))

	health.NewHealthRouter(apiVersion, config.Connect)
	login.NewLoginRouter(apiVersion, config.Connect)
	shoppingcart.NewShoppingCartRouter(apiVersion, config.Connect)
	product.NewProductRouter(apiVersion, config.Connect)
}
