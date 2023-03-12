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

func NewHealthApiResponse() healthApiResponse {
	return healthApiResponse{
		Ready: true,
		Time:  time.Now(),
	}
}

func getHealthService(c *gin.Context) {
	http.Ok(c, NewHealthApiResponse())
}
