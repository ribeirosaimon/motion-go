package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

func NewSessionController(engine *gin.Engine) {
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodGet, "/session", NewLoginService().GetLoggedUserSession),
	).Add()
}
