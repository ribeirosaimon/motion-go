package nosqlDomain

import (
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/shopspring/decimal"
)

type ShoppingCart struct {
	Id        uint64              `bson:"id"`
	Owner     sqlDomain.Profile   `bson:"owner"`
	ProfileId uint64              `bson:"profileId"`
	Price     decimal.Decimal     `bson:"price"`
	Companies []sqlDomain.Company `bson:"companies"`
	BasicNoSQL
}

func (s ShoppingCart) GetId() interface{} {
	return s.Id
}
