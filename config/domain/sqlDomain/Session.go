package sqlDomain

import (
	"time"
)

type Session struct {
	Id        string    `json:"id" gorm:"primary_key"`
	ProfileId uint64    `json:"profileId"`
	LastLogin time.Time `json:"lastLogin"`
}

func (s Session) GetId() interface{} {
	return s.Id
}
