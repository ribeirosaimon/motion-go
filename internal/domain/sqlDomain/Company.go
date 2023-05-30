package sqlDomain

type Company struct {
	// gorm.Model
	Id    uint64 `json:"id" gorm:"primary_key" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Image string `json:"image" bson:"image"`
	BasicSQL
}

func (p Company) GetId() interface{} {
	return p.Id
}
