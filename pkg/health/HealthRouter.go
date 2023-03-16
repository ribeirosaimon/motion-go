package health

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewHealthRouter(engine *gin.Engine, conn func() (*gorm.DB, *sql.DB)) {
	service := NewHealthService()
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodGet, "/health", NewHealthController(service).closeHealth,
			security.Authorization(conn, domain.Role{Name: domain.USER}, domain.Role{Name: domain.ADMIN})),
		controllers.NewMotionRouter(http.MethodGet, "/open-health", NewHealthController(service).openHealth),
	).Add()
}
