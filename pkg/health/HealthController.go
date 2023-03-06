package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
	"github.com/ribeirosaimon/motion-go/pkg/security"
)

func NewHeathController(engine *gin.Engine) {
	controllers.NewMotionController(engine,
		controllers.NewMotionRouter(http.MethodGet, "/health", getHealthService, security.Authorization(security.ADMIN)),
	).Add()
}
