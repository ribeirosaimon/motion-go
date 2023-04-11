package product

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

func NewProductRouter(engine *gin.RouterGroup, conn func() (*gorm.DB, *sql.DB)) {
	service := NewShoppingCartService(conn())
	group := engine.Group("/product")
	config.NewMotionController(group,
		config.NewMotionRouter(http.MethodPost, "", NewProductController(&service).saveProduct,
			security.Authorization(conn, domain.Role{Name: domain.ADMIN})),
	).Add()
}
