package login

import (
	"database/sql"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

func NewLoginRouter(engine *gin.RouterGroup,
	conn func() (*gorm.DB, *sql.DB)) {

	service := NewLoginService(conn())
	controllers.NewMotionController(engine.Group("auth"),
		controllers.NewMotionRouter(http.MethodPost, "/login", NewLoginControler(service).login),
		controllers.NewMotionRouter(http.MethodPost, "/sign-up", NewLoginControler(service).signUp),
	).Add()
}