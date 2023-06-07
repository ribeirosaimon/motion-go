package scraping

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func GetHolders(ticket string) Holders {
	return getHolders(ticket)
}

func getHolders(v string) Holders {
	url := fmt.Sprintf("%s/quote/%s/holders", domain, v)
	c := prepareColly()

	var count = 0
	var holders Holders
	majorHolders := make(map[string]interface{})

	c.OnHTML("table", func(e *colly.HTMLElement) {

		e.ForEach("tbody", func(_ int, el *colly.HTMLElement) {
			if count == 0 {
				e.ForEach("tr", func(_ int, el *colly.HTMLElement) {

					e.ForEach("td", func(v int, el *colly.HTMLElement) {
						if v == 0 {
							majorHolders["allInsiders"] = transformToFloat(el.Text)
						} else if v == 2 {
							majorHolders["sharesInstitutions"] = transformToFloat(el.Text)
						} else if v == 4 {
							majorHolders["heldInstitutions"] = transformToFloat(el.Text)
						} else if v == 6 {
							majorHolders["numberOfInstitutions"] = transformToInteger(el.Text)
						}
					})
				})
				holders.MajorHolders = majorHolders
			} else {
				e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
					var iHolder InstitutionalHolders
					e.ForEach("td", func(y int, el *colly.HTMLElement) {

						if y == 0 {
							iHolder.Holder = el.Text
						} else if y == 1 {
							iHolder.Shares = uint32(transformToInteger(el.Text))
						} else if y == 2 {
							date, _ := TransformDate(el.Text)
							iHolder.DateReported = date
						} else if y == 3 {
							iHolder.PercentOut = uint8(transformToFloat(el.Text))
						} else if y == 4 {
							iHolder.Value = uint64(transformToInteger(el.Text))
						}
					})
					if count == 1 {
						holders.TopInstitutionalHolders = append(holders.TopInstitutionalHolders, iHolder)
					} else {
						holders.TopMutualFundHolders = append(holders.TopMutualFundHolders, iHolder)
					}
				})
			}
			count++
		})
	})
	c.Visit(url)
	c.Wait()
	return holders
}

type Holders struct {
	MajorHolders            map[string]interface{} `json:"majorHolders"`
	TopInstitutionalHolders []InstitutionalHolders `json:"topInstitutionalHolders"`
	TopMutualFundHolders    []InstitutionalHolders `json:"TopMutualFundHolders"`
}

type InstitutionalHolders struct {
	Holder       string    `json:"holders"`
	Shares       uint32    `json:"shares"`
	DateReported time.Time `json:"DateReported"`
	PercentOut   uint8     `json:"percentOut"`
	Value        uint64    `json:"value"`
}
