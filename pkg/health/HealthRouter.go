package health

import (
	"database/sql"
	"net/http"

	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewHealthRouter(conn func() (*gorm.DB, *sql.DB)) config.MotionController {
	service := NewHealthService()
	return config.NewMotionController(
		"/health",
		config.NewMotionRouter(http.MethodGet, "/close", NewHealthController(&service).closeHealth,
			security.Authorization(conn, sqlDomain.Role{Name: sqlDomain.USER}, sqlDomain.Role{Name: sqlDomain.ADMIN})),
		config.NewMotionRouter(http.MethodGet, "/open", NewHealthController(&service).openHealth),
	)
}
