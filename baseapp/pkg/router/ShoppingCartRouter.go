package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/controller"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewShoppingCartRouter() config.MotionController {

	return config.NewMotionController(
		"/shopping-cart",
		config.NewMotionRouter(http.MethodPost, "/create", controller.NewShoppingCartController().CreateShoppingCart,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "", controller.NewShoppingCartController().GetShoppingCart,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodDelete, "", controller.NewShoppingCartController().ExcludeShoppingCart,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodPost, "/company", controller.NewShoppingCartController().AddProductInShoppingCart,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
	)
}
