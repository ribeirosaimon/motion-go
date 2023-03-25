package domain

import (
	"time"
)

type MotionUser struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"-"`
	LastName    string    `json:"lastName,omitempty"`
	Birthday    time.Time `json:"bithday,omitempty"`
	LastLogin   time.Time `json:"lastLogin,omitempty"`
	LoginCount  uint64    `json:"loginCount"`
	LoginAttemp uint8     `json:"loginAttemp,omitempty"`
	BasicSQL
}

func (user MotionUser) GetId() interface{} {
	return user.Id
}
