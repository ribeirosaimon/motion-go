package health

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/config/http"
	"github.com/ribeirosaimon/motion-go/pkg/security"
)

type healthApiResponse struct {
	Ready      bool                `json:"ready"`
	Time       time.Time           `json:"time"`
	LoggedUSer security.LoggedUser `json:"loggedUser"`
}

func getOpenHealthService(c *gin.Context) {
	var response = healthApiResponse{
		Ready: true,
		Time:  time.Now(),
	}
	http.Ok(c, response)
}

func getHealthService(c *gin.Context) {
	var response = healthApiResponse{
		Ready:      true,
		Time:       time.Now(),
		LoggedUSer: security.GetLoggedUser(c),
	}
	http.Ok(c, response)
}
