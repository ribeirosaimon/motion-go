package main

import (
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/shoppingcart/pkg/shoppingcart"
)

func main() {

	motionGo := config.NewMotionGo("config.properties")

	db.Conn = &db.Connections{}
	db.Conn.InitializeDatabases("config.properties")

	motionGo.AddRouter(version1)
	motionGo.RunEngine(motionGo.PropertiesFile.GetInt("server.port", 8081))
}

var version1 = config.RoutersVersion{
	Version: "v1",
	Handlers: []func() config.MotionController{
		shoppingcart.NewShoppingCartRouter,
	},
}
