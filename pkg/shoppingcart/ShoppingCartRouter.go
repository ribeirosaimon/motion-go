package shoppingcart

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewShoppingCartRouter(engine *gin.RouterGroup, conn func() (*gorm.DB, *sql.DB)) {
	service := NewShoppingCartService(conn())
	group := engine.Group("/shopping-cart")
	controllers.NewMotionController(group,
		controllers.NewMotionRouter(http.MethodPost, "/create", NewShoppingCartController(service).createShoppingCart,
			security.Authorization(conn, domain.Role{Name: domain.USER})),
		controllers.NewMotionRouter(http.MethodGet, "", NewShoppingCartController(service).getShoppingCart,
			security.Authorization(conn, domain.Role{Name: domain.USER})),
	).Add()
}
