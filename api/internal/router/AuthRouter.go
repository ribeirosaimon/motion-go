package router

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/api/internal/config"
	"github.com/ribeirosaimon/motion-go/api/internal/controller"
	"github.com/ribeirosaimon/motion-go/api/internal/middleware"
)

func NewLoginRouter() config.MotionController {

	return config.NewMotionController(
		"/auth",
		config.NewMotionRouter(http.MethodPost, "/login", controller.NewAuthController().Login),
		config.NewMotionRouter(http.MethodPost, "/sign-up", controller.NewAuthController().SignUp),
		config.NewMotionRouter(http.MethodGet, "/whoami", controller.NewAuthController().WhoAmI,
			middleware.Authorization()),
		config.NewMotionRouter(http.MethodPost, "/validate", controller.NewAuthController().ValidateEmail,
			middleware.Authorization()),
	)
}
