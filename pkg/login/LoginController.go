package login

import (
	"database/sql"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

func NewLoginController(engine *gin.Engine,
	conn func() (*gorm.DB, *sql.DB)) {
	connect, closeDb := conn()
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodPost, "/login", NewLoginService(connect, closeDb).loginUserService),
		controllers.NewMotionRouter(http.MethodPost, "/sign-up", NewLoginService(connect, closeDb).signUpService),
	).Add()
}
