package domain

import "github.com/shopspring/decimal"

type ShoppingCart struct {
	Id        uint64          `json:"id"`
	Owner     Profile         `json:"owner" gorm:"foreignkey:Id"`
	ProfileId uint64          `json:"profileId"`
	Price     decimal.Decimal `json:"price"`
	Products  []Product       `json:"products" gorm:"many2many:cart_product;"`
	BasicSQL
}

func (s ShoppingCart) GetId() interface{} {
	return s.Id
}
