package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

type userController struct {
	userService userService
}

func NewUserController(engine *gin.Engine) {
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodGet, "/user", newUserService().getUserService),
		controllers.NewMotionRouter(http.MethodPost, "/user", newUserService().saveUserService),
		controllers.NewMotionRouter(http.MethodPut, "/user", newUserService().updateUserService),
		controllers.NewMotionRouter(http.MethodDelete, "/user", newUserService().deleteUser),
	).Add()
}
