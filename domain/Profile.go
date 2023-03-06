package domain

import (
	"time"

	"github.com/ribeirosaimon/motion-go/pkg/security"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	id         uint64              `json:"id"`
	name       string              `json:"name"`
	familyName string              `json:"familyName"`
	age        uint8               `json:"age"`
	birthday   time.Time           `json:"birthday"`
	status     Status              `json:"status"`
	roles      []security.RoleEnum `json:"roles"`
	createdAt  time.Time           `json:"createdAt"`
	updatedAt  time.Time           `json:"updatedAt"`
}

type Status string

const (
	ACTIVE   Status = "ACTIVE"
	INACTIVE        = "INACTIVE"
	BANISH          = "BANISH"
)
