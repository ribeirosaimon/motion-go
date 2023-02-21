package health

import (
	"net/http"

	"github.com/ribeirosaimon/motion-go/pkg/config/controllers"
)

func NewHeathController() {
	controllers.NewMotionController(
		controllers.NewMotionRouter(http.MethodGet, "/health", getHealthService),
	).Add()
}
