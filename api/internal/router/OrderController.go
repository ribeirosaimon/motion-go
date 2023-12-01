package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/api/internal/config"
	"github.com/ribeirosaimon/motion-go/api/internal/controller"
	"github.com/ribeirosaimon/motion-go/api/internal/middleware"
	"github.com/ribeirosaimon/motion-go/confighub/domain/sqlDomain"
)

func NewOrderRouter() config.MotionController {

	return config.NewMotionController(
		"/order",
		config.NewMotionRouter(http.MethodPost, "/save", controller.NewOrderController().NewOrder,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER})),
		config.NewMotionRouter(http.MethodGet, "", controller.NewOrderController().FindAll,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER}, sqlDomain.Role{Name: sqlDomain.ADMIN})),
	)

}
