package health

import (
	"time"

	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

type healthApiResponse struct {
	Ready      bool                  `json:"ready"`
	Time       time.Time             `json:"time"`
	LoggedUser middleware.LoggedUser `json:"loggedUser"`
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

func (s healthService) getHealthService(loggedUser middleware.LoggedUser) healthApiResponse {
	return healthApiResponse{
		Ready:      true,
		Time:       time.Now(),
		LoggedUser: loggedUser,
	}
}
