package product

import (
	"database/sql"
	"net/http"

	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewProductRouter(conn func() (*gorm.DB, *sql.DB)) config.MotionController {
	service := NewProductService(conn())
	return config.NewMotionController(
		"/product",
		config.NewMotionRouter(http.MethodGet, "/:productId", NewProductController(&service).getProduct,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodPost, "", NewProductController(&service).saveProduct,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodPut, "/:productId", NewProductController(&service).updateProduct,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodDelete, "/:productId", NewProductController(&service).deleteProduct,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.ADMIN})),
	)
}
