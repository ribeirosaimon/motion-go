package shoppingcart

import (
	"github.com/ribeirosaimon/motion-go/internal/db"
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/security"
)

func NewShoppingCartRouter(conn *db.Connections) config.MotionController {
	service := NewShoppingCartService(conn)

	return config.NewMotionController(
		"/shopping-cart",
		config.NewMotionRouter(http.MethodPost, "/create", NewShoppingCartController(&service).createShoppingCart,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "", NewShoppingCartController(&service).getShoppingCart,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodDelete, "", NewShoppingCartController(&service).excludeShoppingCart,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodPost, "/product", NewShoppingCartController(&service).addProductInShoppingCart,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.USER})),
	)
}
