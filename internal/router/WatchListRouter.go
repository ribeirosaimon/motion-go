package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/controller"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewPortfolioRouter() config.MotionController {

	return config.NewMotionController(
		"/watchlist",
		config.NewMotionRouter(http.MethodPost, "", controller.NewWatchListController().CreateWatchList,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "", controller.NewWatchListController().GetWatchList,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodDelete, "", controller.NewWatchListController().ExcludeWatchList,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodPost, "/company/:id", controller.NewWatchListController().AddCompanyInWatchList,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodPost, "/company/code/:companyCode", controller.NewWatchListController().
			AddCompanyByCodeInWatchList,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
	)
}
