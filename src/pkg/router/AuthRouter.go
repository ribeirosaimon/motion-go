package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/src/pkg/controller"
)

func NewLoginRouter() config.MotionController {

	return config.NewMotionController(
		"/auth",
		config.NewMotionRouter(http.MethodPost, "/login", controller.NewAuthController().Login),
		config.NewMotionRouter(http.MethodPost, "/sign-up", controller.NewAuthController().SignUp),
		config.NewMotionRouter(http.MethodGet, "/whoami", controller.NewAuthController().WhoAmI,
			middleware.Authorization()),
	)
}
