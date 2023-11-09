package nosqlDomain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WatchList struct {
	BasicNoSQL `bson:"inline"`
	Id         primitive.ObjectID   `json:"id" bson:"_id" gorm:"primaryKey"`
	OwnerId    uint64               `json:"ownerId" bson:"ownerId"`
	Companies  []primitive.ObjectID `json:"companies" bson:"companies"`
}

func (s WatchList) GetId() interface{} {
	return s.Id.Hex()
}
