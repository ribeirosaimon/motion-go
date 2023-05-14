package health

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/security"
)

func NewHealthRouter() config.MotionController {

	service := NewHealthService()
	return config.NewMotionController(
		"/health",
		config.NewMotionRouter(http.MethodGet, "/close", NewHealthController(&service).closeHealth,
			security.Authorization(sqlDomain.Role{Name: sqlDomain.USER}, sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodGet, "/open", NewHealthController(&service).openHealth),
	)
}
