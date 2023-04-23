package login

import (
	"database/sql"
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(engine *gin.RouterGroup,
	conn func() (*gorm.DB, *sql.DB)) {

	service := NewLoginService(conn())
	config.NewMotionController(engine.Group("auth"),
		config.NewMotionRouter(http.MethodPost, "/login", NewLoginControler(&service).login),
		config.NewMotionRouter(http.MethodPost, "/sign-up", NewLoginControler(&service).signUp),
	).Add()
}
