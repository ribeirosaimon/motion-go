package dto

import (
	"time"

	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

type HealthApiResponseDTO struct {
	Ready      bool                  `json:"ready"`
	Time       time.Time             `json:"time"`
	LoggedUser middleware.LoggedUser `json:"loggedUser"`
}
