package nosqlDomain

import (
	"time"

	"github.com/ribeirosaimon/motion-go/confighub/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SummaryStock struct {
	Id          primitive.ObjectID `json:"id"  bson:"_id" gorm:"primary_key"`
	CompanyName string             `json:"companyName" bson:"companyName"`
	CompanyCode string             `json:"companyCode" bson:"companyCode"`
	IsNational  bool               `json:"isNational" bson:"isNational"`
	StockValue  SumarryStockValue  `json:"stockValue" bson:"stockValue"`
	Summary     Summary            `json:"summary" bson:"summary"`
	Status      SummaryStatus      `json:"status" bson:"status"`
	BasicNoSQL  `bson:"inline"`
}

func (s SummaryStock) GetId() interface{} {
	return s.Id.Hex()
}

type SumarryStockValue struct {
	Price        float64 `json:"price" bson:"price"`
	RangeDay     float64 `json:"rangeDay" bson:"rangeDay"`
	PercentRange float64 `json:"percentRange" bson:"percentRange"`
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

type SummaryStatus string

const (
	ACTIVE   SummaryStatus = "ACTIVE"
	INACTIVE SummaryStatus = "INACTIVE"
)

func ChangeProtoToMongo(protoDomain pb.SummaryStock) *SummaryStock {
	status := protoDomain.Status
	summaryStatus := SummaryStatus(status.String())
	summary := createSummary(*protoDomain.GetSummary())
	summaryStockValue := craeteSummaryStockValue(*protoDomain.StockValue)
	hex, err := primitive.ObjectIDFromHex(protoDomain.Id)
	var objectId primitive.ObjectID
	if err != nil {
		objectId = hex
	} else {
		objectId = primitive.NewObjectID()
	}
	return &SummaryStock{
		Id:          objectId,
		Summary:     *summary,
		Status:      summaryStatus,
		StockValue:  *summaryStockValue,
		BasicNoSQL:  BasicNoSQL{CreatedAt: time.Now()},
		CompanyCode: protoDomain.CompanyCode,
		CompanyName: protoDomain.CompanyName,
	}
}

func craeteSummaryStockValue(protoValue pb.SumarryStockValue) *SumarryStockValue {
	return &SumarryStockValue{
		Price:        protoValue.Price,
		RangeDay:     protoValue.RangeDay,
		PercentRange: protoValue.PercentRange,
	}
}

func createSummary(protoSum pb.Summary) *Summary {

	var rangeDayPrice = RangePrice{
		Start: protoSum.GetDayRange().GetStart(),
		End:   protoSum.GetDayRange().GetEnd(),
	}

	var rangeYearPrice = RangePrice{
		Start: protoSum.GetYearRange().GetStart(),
		End:   protoSum.GetYearRange().GetEnd(),
	}

	return &Summary{
		AvgVol:        protoSum.GetAvgVol(),
		Volume:        protoSum.GetVolume(),
		Open:          protoSum.GetOpen(),
		PreviousClose: protoSum.GetPreviousClose(),
		DayRange:      rangeDayPrice,
		YearRange:     rangeYearPrice,
	}
}
