package shoppingcart

import (
	"database/sql"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"

	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewShoppingCartRouter(conn func() (*gorm.DB, *sql.DB), client *mongo.Client) config.MotionController {
	db, s := conn()
	service := NewShoppingCartService(db, s, client)

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
