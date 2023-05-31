package sqlDomain

import (
	"time"

	"github.com/ribeirosaimon/motion-go/internal/domain"
)

type BasicSQL struct {
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
	Status    domain.Status `json:"status" bson:"status"`
}
