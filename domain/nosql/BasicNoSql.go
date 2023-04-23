package nosql

import (
	"time"

	"github.com/ribeirosaimon/motion-go/domain"
)

type BasicNoSQL struct {
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Status    domain.Status `json:"status"`
}
