package domain

import (
	"time"

	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	id        uint64            `json:"id"`
	profileId Profile           `json:"profileId"`
	role      security.RoleEnum `json:"role"`
	lastLogin time.Time         `json:"lastLogin"`
}
