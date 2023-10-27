package router

import (
	"github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/controller"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewHealthRouter() config.MotionController {
	return config.NewMotionController(
		"/health",
		config.NewMotionRouter(http.MethodGet, "/close", controller.NewHealthController().CloseHealth,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER}, sqlDomain.Role{Name: sqlDomain.ADMIN}),
		),
		config.NewMotionRouter(http.MethodGet, "/open", controller.NewHealthController().OpenHealth),
		config.NewMotionRouter(http.MethodGet, "/config", controller.NewHealthController().GetConfigValue),
	)
}
