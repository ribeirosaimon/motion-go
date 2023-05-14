package login

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
)

func NewLoginRouter() config.MotionController {

	service := NewLoginService(db.Conn)
	return config.NewMotionController(
		"/auth",
		config.NewMotionRouter(http.MethodPost, "/login", NewLoginControler(&service).login),
		config.NewMotionRouter(http.MethodPost, "/sign-up", NewLoginControler(&service).signUp),
	)
}
