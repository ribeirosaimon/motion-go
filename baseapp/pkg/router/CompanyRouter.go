package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/controller"
	"github.com/ribeirosaimon/motion-go/internal/middleware"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
)

func NewCompanyRouter() config.MotionController {

	return config.NewMotionController(
		"/company",
		config.NewMotionRouter(http.MethodGet, "/:id", controller.NewCompanyController().GetProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodGet, "/:companyName", controller.NewCompanyController().getCompanyInfo,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodPost, "", controller.NewCompanyController().SaveCompany,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodPut, "/:id", controller.NewCompanyController().UpdateProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodDelete, "/:id", controller.NewCompanyController().DeleteProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
	)
}
