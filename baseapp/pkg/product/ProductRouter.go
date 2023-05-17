package product

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/middleware"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
)

func NewProductRouter() config.MotionController {
	service := NewProductService(db.Conn)
	return config.NewMotionController(
		"/product",
		config.NewMotionRouter(http.MethodGet, "/:id", NewProductController(&service).getProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodPost, "", NewProductController(&service).saveProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodPut, "/:id", NewProductController(&service).updateProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodDelete, "/:id", NewProductController(&service).deleteProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
	)
}
