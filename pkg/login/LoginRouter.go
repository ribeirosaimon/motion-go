package login

import (
	"database/sql"
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"gorm.io/gorm"
)

func NewLoginRouter(
	conn func() (*gorm.DB, *sql.DB)) config.MotionController {

	service := NewLoginService(conn())
	return config.NewMotionController(
		"/auth",
		config.NewMotionRouter(http.MethodPost, "/login", NewLoginControler(&service).login),
		config.NewMotionRouter(http.MethodPost, "/sign-up", NewLoginControler(&service).signUp),
	)
}
