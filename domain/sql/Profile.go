package sql

import (
	"time"
)

type Profile struct {
	Id         uint64     `json:"id" gorm:"primary_key"`
	Name       string     `json:"name"`
	FamilyName string     `json:"familyName"`
	Age        uint8      `json:"age"`
	Birthday   time.Time  `json:"birthday"`
	UserId     uint64     `json:"userId"`
	User       MotionUser `json:"user" gorm:"foreignkey:Id"`
	Roles      []Role     `json:"roles" gorm:"many2many:profile_roles;"`
	BasicSQL
}

func (p Profile) HaveRole(role RoleEnum) bool {
	for _, a := range p.Roles {
		if a.Name == role {
			return true
		}
	}
	return false
}

func (p Profile) GetId() interface{} {
	return p.Id
}

type RoleList []string