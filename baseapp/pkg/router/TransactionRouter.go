package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/controller"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewTransactionRouter() config.MotionController {

	return config.NewMotionController(
		"/transaction",
		config.NewMotionRouter(http.MethodPost, "/buy", controller.NewPortfolioController().CreatePortfolio,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})))
}
