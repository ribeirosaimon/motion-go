package scraping

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/shopspring/decimal"
)

func getStockOptions(v string) StockOptions {
	url := fmt.Sprintf("%s/quote/%s/options", domain, v)
	c := prepareColly()

	var stockOptions StockOptions

	var count int
	c.OnHTML("table", func(e *colly.HTMLElement) {

		e.ForEach("tbody", func(_ int, tbody *colly.HTMLElement) {

			tbody.ForEach("tr", func(_ int, tr *colly.HTMLElement) {

				options := getOptionsValue(tr)
				if count == 0 {
					stockOptions.Calls = append(stockOptions.Calls, options)
				} else {
					stockOptions.Puts = append(stockOptions.Puts, options)
				}
			})

			count++
		})
	})
	c.Visit(url)
	c.Wait()
	return stockOptions

}

func getOptionsValue(tr *colly.HTMLElement) Options {
	var options Options
	tr.ForEach("td", func(v int, el *colly.HTMLElement) {
		if v == 0 {
			options.ContractName = el.Text
		} else if v == 1 {
			options.LastTradeDate, _ = TransformDate(el.Text)
		} else if v == 2 {
			options.Strike = float64(transformToFloat(el.Text))
		} else if v == 3 {
			options.LastPrice = TransformToPrice(el.Text)
		} else if v == 4 {
			options.Bid = TransformToPrice(el.Text)
		} else if v == 5 {
			options.Ask = TransformToPrice(el.Text)
		} else if v == 6 {
			options.Change = TransformToPrice(el.Text)
		} else if v == 7 {
			options.PercentChange = transformToFloat(el.Text)
		} else if v == 8 {
			options.Volume = uint64(transformToInteger(el.Text))
		} else if v == 9 {
			options.OpenInterest = uint64(transformToInteger(el.Text))
		} else if v == 10 {
			options.ImpliedVolatility = transformToFloat(el.Text)
		}
	})
	return options
}

type StockOptions struct {
	Puts  []Options `json:"puts"`
	Calls []Options `json:"calls"`
}

type Options struct {
	ContractName      string          `json:"contractName"`
	LastTradeDate     time.Time       `json:"lastTradeDate"`
	Strike            float64         `json:"strike"`
	LastPrice         decimal.Decimal `json:"lastPrice"`
	Bid               decimal.Decimal `json:"bid"`
	Ask               decimal.Decimal `json:"ask"`
	Change            decimal.Decimal `json:"change"`
	PercentChange     float32         `json:"percentChange"`
	Volume            uint64          `json:"volume"`
	OpenInterest      uint64          `json:"openInterest"`
	ImpliedVolatility float32         `json:"impliedVolatility"`
}
