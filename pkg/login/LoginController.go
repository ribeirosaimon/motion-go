package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

func NewLoginController(engine *gin.Engine) {
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodPost, "/login", NewLoginService().loginUserService),
		controllers.NewMotionRouter(http.MethodPost, "/sign-up", NewLoginService().signUpService),
	).Add()
}
