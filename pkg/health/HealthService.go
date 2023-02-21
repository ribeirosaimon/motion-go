package health

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/http"
)

type healthApiResponse struct {
	Ready bool      `json:"ready"`
	Time  time.Time `json:"time"`
}

func getHealthService(c *gin.Context) {
	response := healthApiResponse{
		Ready: true,
		Time:  time.Now(),
	}
	http.Ok(c, response)
}
