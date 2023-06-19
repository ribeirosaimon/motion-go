package nosqlDomain

import (
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
	Price        float64 `json:"price" bson:"price"`
	RangeDay     float64 `json:"rangeDay" bson:"rangeDay"`
	PersentRange float64 `json:"percentRange" bson:"persentRange"`
}

type Summary struct {
	PreviousClose float64    `json:"previousClose" bson:"previousClose"`
	Open          float64    `json:"open" bson:"open"`
	DayRange      RangePrice `json:"dayRange" bson:"dayRange"`
	YearRange     RangePrice `json:"yearRange" bson:"yearRange"`
	Volume        uint64     `json:"volume" bson:"volume"`
	AvgVol        uint64     `json:"avgVol" bson:"avgVol"`
}

type RangePrice struct {
	Start float64 `json:"start" bson:"start"`
	End   float64 `json:"end" bson:"end"`
}
