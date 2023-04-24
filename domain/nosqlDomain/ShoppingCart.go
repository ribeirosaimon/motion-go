package nosqlDomain

import (
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/shopspring/decimal"
)

type ShoppingCart struct {
	Id        uint64              `json:"id"`
	Owner     sqlDomain.Profile   `json:"owner" gorm:"foreignkey:Id"`
	ProfileId uint64              `json:"profileId"`
	Price     decimal.Decimal     `json:"price"`
	Products  []sqlDomain.Product `json:"products" gorm:"many2many:cart_product;"`
	BasicNoSQL
}

func (s ShoppingCart) GetId() interface{} {
	return s.Id
}
