package service

import (
	"time"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/src/pkg/dto"
)

type HealthService struct {
}

func NewHealthService() HealthService {
	return HealthService{}
}

func (s *HealthService) GetOpenHealthService() dto.HealthApiResponseDTO {
	return dto.HealthApiResponseDTO{
		Ready: true,
		Time:  time.Now(),
	}

}

func (s *HealthService) GetHealthService(loggedUser middleware.LoggedUser) dto.HealthApiResponseDTO {
	return dto.HealthApiResponseDTO{
		Ready:      true,
		Time:       time.Now(),
		LoggedUser: loggedUser,
	}
}

func (s *HealthService) GetConfigurations() config.MotionConfig {
	return config.GetConfiguration()
}
