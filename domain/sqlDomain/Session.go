package sqlDomain

import (
	"time"
)

type Session struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	SessionId string    `json:"sessionId" gorm:"foreignkey:Id"`
	ProfileId uint64    `json:"profileId"`
	LastLogin time.Time `json:"lastLogin"`
	BasicSQL
}

func (s Session) GetId() interface{} {
	return s.Id
}
