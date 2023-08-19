package nosqlDomain

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Portfolio struct {
	BasicNoSQL `bson:"inline"`
	Id         primitive.ObjectID `json:"id" bson:"_id" gorm:"primaryKey"`
	OwnerId    uint64             `json:"ownerId" bson:"ownerId"`
	Price      decimal.Decimal    `json:"price" bson:"price"`
	Companies  []MineStock        `json:"companies" bson:"companies"`
}

func (s Portfolio) GetId() interface{} {
	return s.Id.Hex()
}

type MineStock struct {
	StockId  primitive.ObjectID `json:"stockId" bson:"stockId"`
	BuyPrice decimal.Decimal    `json:"buyPrice" bson:"buyPrice"`
	Quantity float64            `json:"quantity" bson:"quantity"`
}

func (m MineStock) CalculeValue() decimal.Decimal {
	return m.BuyPrice.Mul(decimal.NewFromFloat(m.Quantity))
}
