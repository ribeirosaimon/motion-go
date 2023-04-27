package version

import (
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/pkg/login"
	"github.com/ribeirosaimon/motion-go/pkg/product"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"
)

var V1 = config.RoutersVersion{
	Version: "v1",
	Handlers: []config.MotionController{
		health.NewHealthRouter(db.ConnectSqlDb),
		login.NewLoginRouter(db.ConnectSqlDb),
		shoppingcart.NewShoppingCartRouter(db.ConnectSqlDb),
		product.NewProductRouter(db.ConnectSqlDb),
	},
}

var V2 = config.RoutersVersion{
	Version: "v2",
	Handlers: []config.MotionController{
		health.NewHealthRouter(db.ConnectSqlDb),
	},
}
