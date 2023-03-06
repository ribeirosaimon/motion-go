package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

func NewLoginController(engine *gin.Engine) {
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodGet, "/login", loginService),
		controllers.NewMotionRouter(http.MethodGet, "/sign-up", signUpService),
	).Add()
}
