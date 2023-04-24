package sqlDomain

import (
	"time"

	"github.com/ribeirosaimon/motion-go/domain"
)

type BasicSQL struct {
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Status    domain.Status `json:"status"`
}
