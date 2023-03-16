package domain

import "github.com/shopspring/decimal"

type ShoppingCart struct {
	Id       uint64          `json:"id"`
	Products []Product       `json:"products"`
	Price    decimal.Decimal `json:"price"`
}
