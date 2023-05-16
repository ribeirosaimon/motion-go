package health

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func NewHealthRouter() config.MotionController {
	path := "/health"
	service := NewHealthService()

	return config.NewMotionController(
		path,
		config.NewMotionRouter(http.MethodGet, "/close", NewHealthController(&service).closeHealth,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER}, sqlDomain.Role{Name: sqlDomain.ADMIN}),
		),
		config.NewMotionRouter(http.MethodGet, "/open", NewHealthController(&service).openHealth),
	)
}
