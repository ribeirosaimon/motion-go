package health

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	sql2 "github.com/ribeirosaimon/motion-go/domain/sql"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewHealthRouter(engine *gin.RouterGroup, conn func() (*gorm.DB, *sql.DB)) {
	service := NewHealthService()
	group := engine.Group("/health")
	config.NewMotionController(group,
		config.NewMotionRouter(http.MethodGet, "/close", NewHealthController(&service).closeHealth,
			security.Authorization(conn, sql2.Role{Name: sql2.USER}, sql2.Role{Name: sql2.ADMIN})),
		config.NewMotionRouter(http.MethodGet, "/open", NewHealthController(&service).openHealth),
	).Add()
}
