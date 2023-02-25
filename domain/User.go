package domain

import (
	"time"

	"gorm.io/gorm"
)

type MotionUser struct {
	gorm.Model
	Id          uint64    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name,omitempty" `
	Email       string    `json:"email,omitempty" `
	LastName    string    `json:"lastName,omitempty"`
	Birthday    time.Time `json:"bithday,omitempty"`
	LastLogin   time.Time `json:"lastLogin,omitempty"`
	LoginCount  uint64    `json:"loginCount"`
	LoginAttemp uint8     `json:"loginAttemp,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

func (user MotionUser) GetId() interface{} {
	return user.Id
}
