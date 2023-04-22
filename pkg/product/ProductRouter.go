package product

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewProductRouter(engine *gin.RouterGroup, conn func() (*gorm.DB, *sql.DB)) {
	service := NewShoppingCartService(conn())
	group := engine.Group("/product")
	config.NewMotionController(group,
		config.NewMotionRouter(http.MethodGet, "/:productId", NewProductController(&service).getProduct,
			security.Authorization(conn, domain.Role{Name: domain.ADMIN})),
		config.NewMotionRouter(http.MethodPost, "", NewProductController(&service).saveProduct,
			security.Authorization(conn, domain.Role{Name: domain.ADMIN})),
		config.NewMotionRouter(http.MethodPut, "/:productId", NewProductController(&service).updateProduct,
			security.Authorization(conn, domain.Role{Name: domain.ADMIN})),
		config.NewMotionRouter(http.MethodDelete, "/:productId", NewProductController(&service).deleteProduct,
			security.Authorization(conn, domain.Role{Name: domain.ADMIN})),
	).Add()
}
