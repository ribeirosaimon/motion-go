package shoppingcart

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/security"
)

func NewShoppingCartRouter() config.MotionController {
	service := NewShoppingCartService(db.Conn)

	return config.NewMotionController(
		"/shopping-cart",
		config.NewMotionRouter(http.MethodPost, "/create", NewShoppingCartController(&service).createShoppingCart,
			security.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "", NewShoppingCartController(&service).getShoppingCart,
			security.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodDelete, "", NewShoppingCartController(&service).excludeShoppingCart,
			security.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodPost, "/product", NewShoppingCartController(&service).addProductInShoppingCart,
			security.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
	)
}
