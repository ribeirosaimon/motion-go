package router

import (
	"github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/controller"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewTransactionRouter() config.MotionController {

	return config.NewMotionController(
		"/transaction",
		config.NewMotionRouter(http.MethodPost, "/deposit", controller.NewTransactionController().Deposit,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "/balance", controller.NewTransactionController().Balance,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "", controller.NewTransactionController().FindAllTransactions,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER}, sqlDomain.Role{Name: sqlDomain.ADMIN})),
	)

}
