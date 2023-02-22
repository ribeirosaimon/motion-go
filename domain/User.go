package domain

import "time"

type User struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	LastName    string    `json:"lastName,omitempty"`
	Birthday    time.Time `json:"bithday,omitempty"`
	LastLogin   time.Time `json:"lastLogin,omitempty"`
	LoginCount  uint64    `json:"loginCount"`
	LoginAttemp uint8     `json:"loginAttemp,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}
