package health

import (
	"time"

	"github.com/ribeirosaimon/motion-go/internal/security"
)

type healthApiResponse struct {
	Ready      bool                `json:"ready"`
	Time       time.Time           `json:"time"`
	LoggedUser security.LoggedUser `json:"loggedUser"`
}

type healthService struct{}

func NewHealthService() healthService {
	return healthService{}
}

func (s healthService) getOpenHealthService() healthApiResponse {
	return healthApiResponse{
		Ready: true,
		Time:  time.Now(),
	}

}

func (s healthService) getHealthService(loggedUser security.LoggedUser) healthApiResponse {
	return healthApiResponse{
		Ready:      true,
		Time:       time.Now(),
		LoggedUser: loggedUser,
	}
}
