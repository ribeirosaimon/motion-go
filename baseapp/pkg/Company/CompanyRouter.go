package Company

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/middleware"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
)

func NewCompanyRouter() config.MotionController {

	service := NewCompanyService(db.Conn)
	return config.NewMotionController(
		"/company",
		config.NewMotionRouter(http.MethodGet, "/:id", NewCompanyController(&service).getProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodPost, "", NewCompanyController(&service).saveCompany,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodPut, "/:id", NewCompanyController(&service).updateProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodDelete, "/:id", NewCompanyController(&service).deleteProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
	)
}
