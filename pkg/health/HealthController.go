package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

func NewHeathController(engine *gin.Engine) {
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodGet, "/health", getHealthService),
	).Add()
}
