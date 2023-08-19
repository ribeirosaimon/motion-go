package dto

import "github.com/shopspring/decimal"

type BuyPriceDTO struct {
	Price    decimal.Decimal `json:"price"`
	Quantity float64         `json:"quantity"`
}
