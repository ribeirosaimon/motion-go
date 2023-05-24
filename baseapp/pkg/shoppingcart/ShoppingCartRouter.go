package shoppingcart

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewShoppingCartRouter() config.MotionController {
	service := NewShoppingCartService(db.Conn)

	return config.NewMotionController(
		"/shopping-cart",
		config.NewMotionRouter(http.MethodPost, "/create", NewShoppingCartController(&service).createShoppingCart,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "", NewShoppingCartController(&service).getShoppingCart,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodDelete, "", NewShoppingCartController(&service).excludeShoppingCart,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodPost, "/Company", NewShoppingCartController(&service).addProductInShoppingCart,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
	)
}
