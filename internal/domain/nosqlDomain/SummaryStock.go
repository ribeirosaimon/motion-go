package nosqlDomain

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SummaryStock struct {
	Id          primitive.ObjectID `json:"id"  bson:"_id" gorm:"primary_key"`
	CompanyName string             `json:"companyName" bson:"companyName"`
	CompanyCode string             `json:"companyCode" bson:"companyCode"`
	StockValue  SumarryStockValue  `json:"stockValue" bson:"stockValue"`
	Summary     Summary            `json:"summary" bson:"summary"`
	BasicNoSQL  `bson:"inline"`
}

func (s SummaryStock) GetId() interface{} {
	return s.Id.Hex()
}

type SumarryStockValue struct {
	Price        decimal.Decimal `json:"price" bson:"price"`
	RangeDay     decimal.Decimal `json:"rangeDay" bson:"rangeDay"`
	PersentRange float32         `json:"percentRange" bson:"persent_range"`
}

type Summary struct {
	PreviousClose decimal.Decimal `json:"previousClose" json:"previousClose"`
	Open          decimal.Decimal `json:"open" bson:"open"`
	DayRange      RangePrice      `json:"dayRange" bson:"dayRange"`
	YearRange     RangePrice      `json:"yearRange" bson:"yearRange"`
	Volume        uint64          `json:"volume" bson:"volume"`
	AvgVol        uint64          `json:"avgVol" bson:"avgVol"`
}

type RangePrice struct {
	Start decimal.Decimal `json:"start" bson:"start"`
	End   decimal.Decimal `json:"end" bson:"end"`
}
