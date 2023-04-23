package domain

import "time"

type BasicSQL struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status    Status    `json:"status"`
}
