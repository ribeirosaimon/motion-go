package sql

type Role struct {
	Id   uint     `json:"id" gorm:"primary_key"`
	Name RoleEnum `json:"name" gorm:"unique"`
}

func (r Role) GetId() interface{} {
	return r.Id
}

type RoleEnum string

const (
	ADMIN RoleEnum = "ADMIN"
	USER  RoleEnum = "USER"
)
