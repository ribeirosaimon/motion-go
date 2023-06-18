package sqlDomain

type Company struct {
	Id    uint64 `json:"id" bson:"id" gorm:"primary_key"`
	Name  string `json:"name" bson:"name"`
	Image string `json:"image" bson:"image"`
	BasicSQL
}

func (p Company) GetId() interface{} {
	return p.Id
}
