package product

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	sql2 "github.com/ribeirosaimon/motion-go/domain/sql"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewProductRouter(engine *gin.RouterGroup, conn func() (*gorm.DB, *sql.DB)) {
	service := NewProductService(conn())
	group := engine.Group("/product")
	config.NewMotionController(group,
		config.NewMotionRouter(http.MethodGet, "/:productId", NewProductController(&service).getProduct,
			security.Authorization(conn, sql2.Role{Name: sql2.ADMIN})),
		config.NewMotionRouter(http.MethodPost, "", NewProductController(&service).saveProduct,
			security.Authorization(conn, sql2.Role{Name: sql2.ADMIN})),
		config.NewMotionRouter(http.MethodPut, "/:productId", NewProductController(&service).updateProduct,
			security.Authorization(conn, sql2.Role{Name: sql2.ADMIN})),
		config.NewMotionRouter(http.MethodDelete, "/:productId", NewProductController(&service).deleteProduct,
			security.Authorization(conn, sql2.Role{Name: sql2.ADMIN})),
	).Add()
}
