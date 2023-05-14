package product

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/db"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/security"
)

func NewProductRouter() config.MotionController {
	service := NewProductService(db.Conn)
	return config.NewMotionController(
		"/product",
		config.NewMotionRouter(http.MethodGet, "/:productId", NewProductController(&service).getProduct,
			security.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodPost, "", NewProductController(&service).saveProduct,
			security.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodPut, "/:productId", NewProductController(&service).updateProduct,
			security.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodDelete, "/:productId", NewProductController(&service).deleteProduct,
			security.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
	)
}
