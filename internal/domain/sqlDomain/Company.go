package sqlDomain

type Company struct {
	Id    uint64 `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Image string `json:"image"`
	BasicSQL
}

func (p Company) GetId() interface{} {
	return p.Id
}
