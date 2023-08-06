package sqlDomain

import "github.com/shopspring/decimal"

type Transaction struct {
	Id            uint64          `json:"id" gorm:"primary_key"`
	SessionId     string          `json:"sessionId" gorm:"foreignkey:Id"`
	ProfileId     uint64          `json:"profileId"`
	StockId       string          `json:"stockId"`
	Quantity      decimal.Decimal `json:"quantity"`
	Value         decimal.Decimal `json:"value"`
	OperationType OperationType   `json:"operationType"`
	BasicSQL
}

type OperationType string

const (
	BUY  OperationType = "BUY"
	SELL RoleEnum      = "SELL"
)
