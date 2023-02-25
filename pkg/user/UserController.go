package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

func NewUserController(engine *gin.Engine) {
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodGet, "/user", getUserService),
		controllers.NewMotionRouter(http.MethodPost, "/user", saveUserService),
		controllers.NewMotionRouter(http.MethodPut, "/user", updateUserService),
		controllers.NewMotionRouter(http.MethodDelete, "/user", deleteUser),
	).Add()
}
