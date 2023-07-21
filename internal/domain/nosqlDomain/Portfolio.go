package nosqlDomain

import (
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Portfolio struct {
	BasicNoSQL `bson:"inline"`
	Id         primitive.ObjectID   `json:"id" bson:"_id" gorm:"primaryKey"`
	Owner      sqlDomain.Profile    `json:"owner" bson:"owner"`
	ProfileId  uint64               `json:"profileId" bson:"profileId"`
	Price      decimal.Decimal      `json:"price" bson:"price"`
	Companies  []primitive.ObjectID `json:"companies" bson:"companies"`
}

func (s Portfolio) GetId() interface{} {
	return s.Id.Hex()
}
