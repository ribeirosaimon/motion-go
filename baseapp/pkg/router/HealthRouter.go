package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/controller"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewHealthRouter() config.MotionController {
	return config.NewMotionController(
		"/health",
		config.NewMotionRouter(http.MethodGet, "/close", controller.NewHealthController().CloseHealth,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER}, sqlDomain.Role{Name: sqlDomain.ADMIN}),
		),
		config.NewMotionRouter(http.MethodGet, "/open", controller.NewHealthController().OpenHealth),
	)
}
