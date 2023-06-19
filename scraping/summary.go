package scraping

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
)

func GetStockSummary(v string) nosqlDomain.SummaryStock {
	url := fmt.Sprintf("%s/quote/%s", domain, v)
	c := prepareColly()

	var sumStock nosqlDomain.SummaryStock
	c.OnHTML("#Lead-5-QuoteHeader-Proxy", func(e *colly.HTMLElement) {
		e.ForEach("h1", func(_ int, translate *colly.HTMLElement) {
			sumStock.CompanyName = translate.Text
		})
		sumStock.StockValue = getSummaryStockValue(e)
		sumStock.CompanyCode = v
	})

	c.OnHTML("#quote-summary", func(e *colly.HTMLElement) {
		e.ForEach("tbody", func(tbodyNumber int, tbody *colly.HTMLElement) {
			var summary nosqlDomain.Summary
			if tbodyNumber == 0 {
				tbody.ForEach("tr", func(trCount int, tr *colly.HTMLElement) {
					if trCount == 0 {
						getTdValue(tr)
						summary.PreviousClose = transformToFloat(getTdValue(tr))
					} else if trCount == 1 {
						summary.Open = transformToFloat(getTdValue(tr))
					} else if trCount == 4 {
						splitedValue := strings.Split(getTdValue(tr), "-")
						summary.DayRange.Start = transformToFloat(splitedValue[0])
						summary.DayRange.End = transformToFloat(splitedValue[1])
					} else if trCount == 5 {
						splitedValue := strings.Split(getTdValue(tr), "-")
						summary.YearRange.Start = transformToFloat(splitedValue[0])
						summary.YearRange.End = transformToFloat(splitedValue[1])
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

func getSummaryStockValue(v *colly.HTMLElement) nosqlDomain.SumarryStockValue {
	var sumarryStockValue nosqlDomain.SumarryStockValue
	v.ForEach("fin-streamer", func(countValue int, value *colly.HTMLElement) {
		if countValue == 0 {
			sumarryStockValue.Price = transformToFloat(value.Text)
		} else if countValue == 1 {
			sumarryStockValue.RangeDay = transformToFloat(value.Text)
		} else if countValue == 2 {
			sumarryStockValue.PersentRange = transformToFloat(value.Text)
		}
	})
	return sumarryStockValue
}
