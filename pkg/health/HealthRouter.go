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

func NewHealthRouter(engine *gin.RouterGroup, conn func() (*gorm.DB, *sql.DB)) {
	service := NewHealthService()
	group := engine.Group("/health")
	controllers.NewMotionController(group,
		controllers.NewMotionRouter(http.MethodGet, "/close", NewHealthController(service).closeHealth,
			security.Authorization(conn, domain.Role{Name: domain.USER}, domain.Role{Name: domain.ADMIN})),
		controllers.NewMotionRouter(http.MethodGet, "/open", NewHealthController(service).openHealth),
	).Add()
}
