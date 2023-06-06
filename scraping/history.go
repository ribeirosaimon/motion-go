package scraping

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/shopspring/decimal"
)

func GetStock(stock ...string) []StockHistory {
	init := time.Now()
	var result []StockHistory

	var wg sync.WaitGroup
	wg.Add(len(stock))

	for _, v := range stock {
		go func(stock string) {
			defer wg.Done()
			result = append(result, getHistoryPrice(stock, 30))
		}(v)
	}
	wg.Wait()
	fim := time.Now()
	duration := fim.Sub(init)
	fmt.Println(duration)

	return result
}

func getHistoryPrice(v string, dayRange int) StockHistory {

	url := fmt.Sprintf("%s/quote/%s/history", domain, v)

	c := prepareColly()
	c.OnRequest(func(r *colly.Request) {
		log.Println("get value of:", v)
	})
	var stockHistory StockHistory
	c.OnHTML("tbody", func(e *colly.HTMLElement) {

		stockHistory.Name = v

		e.ForEach("tr", func(v int, el *colly.HTMLElement) {

			if v <= dayRange {
				var day dayHistory
				var cnt = 0
				el.ForEach("td", func(_ int, el *colly.HTMLElement) {
					cnt += 1
					if cnt == 1 {
						date, _ := transformDate(historyDataLayout, el.Text)
						day.Date = date
					} else if cnt == 2 {
						day.Open = TransformToPrice(el.Text)
					} else if cnt == 3 {
						day.High = TransformToPrice(el.Text)
					} else if cnt == 4 {
						day.Low = TransformToPrice(el.Text)
					} else if cnt == 5 {
						day.Close = TransformToPrice(el.Text)
					} else if cnt == 6 {
						day.AdjClose = TransformToPrice(el.Text)
					} else if cnt == 7 {
						day.Volume = transformToInteger(el.Text)
					}

				})
				stockHistory.History = append(stockHistory.History, day)
			} else {
				return
			}
			v++
		})

	})

	c.Visit(url)
	c.Wait()
	return stockHistory
}

type StockHistory struct {
	Name    string       `json:"name"`
	History []dayHistory `json:"history"`
}

type dayHistory struct {
	Date     time.Time       `json:"date"`
	Open     decimal.Decimal `json:"open"`
	High     decimal.Decimal `json:"high"`
	Low      decimal.Decimal `json:"low"`
	Close    decimal.Decimal `json:"close"`
	AdjClose decimal.Decimal `json:"adjClose"`
	Volume   int             `json:"volume"`
}
