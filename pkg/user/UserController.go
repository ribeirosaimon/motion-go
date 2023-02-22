package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

func NewUserController(engine *gin.Engine) {
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodGet, "/user", getUserService),
	).Add()
}
