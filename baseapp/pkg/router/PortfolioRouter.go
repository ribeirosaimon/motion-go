package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/controller"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewPortfolioRouter() config.MotionController {

	return config.NewMotionController(
		"/portfolio",
		config.NewMotionRouter(http.MethodPost, "", controller.NewPortfolioController().CreatePortfolio,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "", controller.NewPortfolioController().GetPortfolio,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodDelete, "", controller.NewPortfolioController().ExcludePortfolio,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodPost, "/company/:id", controller.NewPortfolioController().AddCompanyInPortfolio,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
	)
}
