package sqlDomain

type Profile struct {
	Id           uint64        `json:"id" gorm:"primary_key"`
	Name         string        `json:"name"`
	MotionUserId uint64        `json:"-"  gorm:"foreignkey:Id"`
	Code         string        `json:"-"`
	Status       ProfileStatus `json:"status"`
	Roles        []Role        `json:"roles,omitempty" gorm:"many2many:profile_roles;"`
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

type ProfileStatus string

const (
	ACTIVE     ProfileStatus = "ACTIVE"
	INACTIVE   ProfileStatus = "INACTIVE"
	EMAIL_SYNC ProfileStatus = "EMAIL_SYNC"
)
