package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/controller"
	"github.com/ribeirosaimon/motion-go/internal/config"
)

func NewLoginRouter() config.MotionController {

	return config.NewMotionController(
		"/auth",
		config.NewMotionRouter(http.MethodPost, "/login", controller.NewLoginControler().Login),
		config.NewMotionRouter(http.MethodPost, "/sign-up", controller.NewLoginControler().SignUp),
	)
}
