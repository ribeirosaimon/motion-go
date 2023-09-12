package nosqlDomain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Portfolio struct {
	BasicNoSQL `bson:"inline"`
	Id         primitive.ObjectID `json:"id" bson:"_id" gorm:"primaryKey"`
	OwnerId    uint64             `json:"ownerId" bson:"ownerId"`
	Cash       float64            `json:"cash" bson:"cash"`
	Price      float64            `json:"price" bson:"price"`
	Companies  []MineStock        `json:"companies" bson:"companies"`
}

func (s Portfolio) GetId() interface{} {
	return s.Id.Hex()
}

type MineStock struct {
	StockId  primitive.ObjectID `json:"stockId" bson:"stockId"`
	BuyPrice float64            `json:"buyPrice" bson:"buyPrice"`
	Quantity float64            `json:"quantity" bson:"quantity"`
}

func (m MineStock) CalculeValue() float64 {
	return m.BuyPrice * m.Quantity
}
