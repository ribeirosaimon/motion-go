package shoppingcart

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewShoppingCartRouter(engine *gin.RouterGroup, conn func() (*gorm.DB, *sql.DB)) {
	service := NewShoppingCartService(conn())
	group := engine.Group("/shopping-cart")
	config.NewMotionController(group,
		config.NewMotionRouter(http.MethodPost, "/create", NewShoppingCartController(&service).createShoppingCart,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "", NewShoppingCartController(&service).getShoppingCart,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodDelete, "", NewShoppingCartController(&service).excludeShoppingCart,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodPost, "/product", NewShoppingCartController(&service).addProductInShoppingCart,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.USER})),
	).Add()
}
