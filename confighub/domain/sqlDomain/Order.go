package sqlDomain

import (
	"time"
)

type Order struct {
	Id               uint64      `json:"id" gorm:"primary_key"`
	SessionId        string      `json:"-" gorm:"foreignkey:SessionID"`
	ProfileId        uint64      `json:"-" gorm:"foreignkey:ProfileID"`
	Value            float64     `json:"value"`
	Quantity         float64     `json:"quantity"`
	SummaryStockId   string      `json:"summaryStockId"`
	SummaryStockCode string      `json:"summaryStockCode"`
	Status           OrderStatus `json:"status"`
	BasicSQL
}

func (o Order) GetId() interface{} {
	return o.Id
}

type OrderStatus string

const (
	OrderActive   OrderStatus = "ACTIVE"
	OrderCanceled OrderStatus = "CANCELED"
	OrderExecuted OrderStatus = "EXECUTED"
)

func NewOrder(value, quantity float64, profileId uint64, sessionId, summaryStock, summaryStockName string,
	status OrderStatus) Order {

	now := time.Now()
	basicSql := BasicSQL{CreatedAt: now, UpdatedAt: now}

	return Order{
		SessionId:        sessionId,
		ProfileId:        profileId,
		Value:            value,
		Quantity:         quantity,
		SummaryStockId:   summaryStock,
		SummaryStockCode: summaryStockName,
		Status:           status,
		BasicSQL:         basicSql,
	}
}
