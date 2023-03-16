package domain

import "github.com/shopspring/decimal"

type Product struct {
	Id    uint64          `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
	Image string          `json:"image"`
}
