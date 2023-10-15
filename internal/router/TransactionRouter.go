package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/controller"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewTransactionRouter() config.MotionController {

	return config.NewMotionController(
		"/transaction",
		config.NewMotionRouter(http.MethodPost, "/deposit", controller.NewTransactionController().Deposit,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})))
}
