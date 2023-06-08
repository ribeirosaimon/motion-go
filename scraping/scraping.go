package scraping

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func GetStockInformations(v string) StockInfo {
	now := time.Now()
	var stockInfo StockInfo
	stockInfo.Financials = financials(v)
	stockInfo.HistoryPrice = getHistoryPrice(v, 31)
	stockInfo.Holders = getHolders(v)
	stockInfo.Options = getStockOptions(v)
	end := time.Now()
	fmt.Printf("time to scraping this stock: %s \n", end.Sub(now))
	return stockInfo
}

func prepareColly() *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"),
		colly.AllowedDomains("finance.yahoo.com"),
		colly.MaxBodySize(0),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
	})

	return c
}

type StockInfo struct {
	Financials   map[string]map[string]interface{} `json:"financials"`
	HistoryPrice StockHistory                      `json:"historyPrice"`
	Holders      Holders                           `json:"holders"`
	Options      StockOptions                      `json:"options"`
}
