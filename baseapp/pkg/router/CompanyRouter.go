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
		config.NewMotionRouter(http.MethodGet, "/:id", controller.NewCompanyController().GetCompany,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN}, sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "/code/:companyName", controller.NewCompanyController().GetCompanyInfo,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodDelete, "/:id", controller.NewCompanyController().DeleteProduct,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.ADMIN})),
	)
}
