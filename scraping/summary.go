package scraping

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/shopspring/decimal"
)

func GetStockSummary(v string) SummaryStock {
	url := fmt.Sprintf("%s/quote/%s", domain, v)
	c := prepareColly()

	var sumStock SummaryStock
	c.OnHTML("#Lead-5-QuoteHeader-Proxy", func(e *colly.HTMLElement) {
		e.ForEach("h1", func(_ int, translate *colly.HTMLElement) {
			sumStock.CompanyName = translate.Text
		})
		sumStock.StockValue = getSummaryStockValue(e)

	})

	c.OnHTML("#quote-summary", func(e *colly.HTMLElement) {
		e.ForEach("tbody", func(tbodyNumber int, tbody *colly.HTMLElement) {
			var summary Summary
			if tbodyNumber == 0 {
				tbody.ForEach("tr", func(trCount int, tr *colly.HTMLElement) {
					if trCount == 0 {
						getTdValue(tr)
						summary.PreviousClose = TransformToPrice(getTdValue(tr))
					} else if trCount == 1 {
						summary.Open = TransformToPrice(getTdValue(tr))
					} else if trCount == 4 {
						splitedValue := strings.Split(getTdValue(tr), "-")
						summary.DayRange.Start = TransformToPrice(splitedValue[0])
						summary.DayRange.End = TransformToPrice(splitedValue[1])
					} else if trCount == 5 {
						splitedValue := strings.Split(getTdValue(tr), "-")
						summary.YearRange.Start = TransformToPrice(splitedValue[0])
						summary.YearRange.End = TransformToPrice(splitedValue[1])
					} else if trCount == 6 {
						summary.Volume = uint64(transformToInteger(getTdValue(tr)))
					} else if trCount == 7 {
						summary.AvgVol = uint64(transformToInteger(getTdValue(tr)))
					}
				})
				sumStock.Summary = summary
			}
		})

	})
	c.Visit(url)
	c.Wait()
	return sumStock
}

func getTdValue(tr *colly.HTMLElement) string {
	var s string
	tr.ForEach("td", func(tdCount int, td *colly.HTMLElement) {
		if tdCount == 1 {
			s = td.Text
		}
	})
	return s
}

func getSummaryStockValue(v *colly.HTMLElement) SumarryStockValue {
	var sumarryStockValue SumarryStockValue
	v.ForEach("fin-streamer", func(countValue int, value *colly.HTMLElement) {
		if countValue == 0 {
			sumarryStockValue.Price = TransformToPrice(value.Text)
		} else if countValue == 1 {
			sumarryStockValue.RangeDay = TransformToPrice(value.Text)
		} else if countValue == 2 {
			sumarryStockValue.PersentRange = transformToFloat(value.Text)
		}
	})
	return sumarryStockValue
}

type SummaryStock struct {
	CompanyName string            `json:"companyName"`
	StockValue  SumarryStockValue `json:"stockValue"`
	Summary     Summary           `json:"summary"`
}

type SumarryStockValue struct {
	Price        decimal.Decimal `json:"price"`
	RangeDay     decimal.Decimal `json:"rangeDay"`
	PersentRange float32         `json:"percentRange"`
}

type Summary struct {
	PreviousClose decimal.Decimal `json:"previousClose"`
	Open          decimal.Decimal `json:"open"`
	DayRange      RangePrice      `json:"dayRange"`
	YearRange     RangePrice      `json:"yearRange"`
	Volume        uint64          `json:"volume"`
	AvgVol        uint64          `json:"avgVol"`
}

type RangePrice struct {
	Start decimal.Decimal `json:"start"`
	End   decimal.Decimal `json:"end"`
}
