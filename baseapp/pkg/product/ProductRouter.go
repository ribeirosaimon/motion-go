package product

import (
	"github.com/ribeirosaimon/motion-go/internal/db"
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/security"
)

func NewProductRouter(conn *db.Connections) config.MotionController {
	service := NewProductService(conn)
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
