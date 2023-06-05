package scraping

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

const (
	domain = "https://finance.yahoo.com"
	layout = "Jan 02, 2006"
)

type StockHistory struct {
	Name    string       `json:"name"`
	History []dayHistory `json:"history"`
}

type dayHistory struct {
	Date     time.Time `json:"date"`
	Open     float32   `json:"open"`
	High     float32   `json:"high"`
	Low      float32   `json:"low"`
	Close    float32   `json:"close"`
	AdjClose float32   `json:"adjClose"`
	Volume   int       `json:"volume"`
}

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
						day.Date = transformDate(el.Text)
					} else if cnt == 2 {
						day.Open = transformToPrice(el.Text)
					} else if cnt == 3 {
						day.High = transformToPrice(el.Text)
					} else if cnt == 4 {
						day.Low = transformToPrice(el.Text)
					} else if cnt == 5 {
						day.Close = transformToPrice(el.Text)
					} else if cnt == 6 {
						day.AdjClose = transformToPrice(el.Text)
					} else if cnt == 7 {
						day.Volume = transformToVolume(el.Text)
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

func transformToPrice(v string) float32 {
	if strings.Contains(v, ",") {
		newValue := strings.ReplaceAll(v, ",", "")
		if s, err := strconv.ParseFloat(newValue, 32); err == nil {
			return float32(s)
		}
	}
	return float32(0)
}

func transformToVolume(v string) int {
	if strings.Contains(v, ",") {
		newValue := strings.ReplaceAll(v, ",", "")
		if s, err := strconv.Atoi(newValue); err == nil {
			return s
		}
	}
	return 0
}

func transformDate(dateString string) time.Time {
	date, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}
	return date
}
