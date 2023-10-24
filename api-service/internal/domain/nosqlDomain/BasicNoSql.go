package nosqlDomain

import (
	"time"
)

type BasicNoSQL struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
