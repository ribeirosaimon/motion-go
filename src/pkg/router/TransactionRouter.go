package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/src/pkg/controller"
)

func NewTransactionRouter() config.MotionController {

	return config.NewMotionController(
		"/transaction",
		config.NewMotionRouter(http.MethodPost, "/deposit", controller.NewPortfolioController().CreatePortfolio,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})))
}
