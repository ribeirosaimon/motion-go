package scraping

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func getHolders(v string) Holders {
	url := fmt.Sprintf("%s/quote/%s/holders", domain, v)
	c := prepareColly()

	var count = 0
	var holders Holders
	majorHolders := make(map[string]interface{})

	c.OnHTML("table", func(e *colly.HTMLElement) {

		e.ForEach("tbody", func(_ int, tr *colly.HTMLElement) {
			if count == 0 {
				tr.ForEach("tr", func(_ int, td *colly.HTMLElement) {

					td.ForEach("td", func(v int, el *colly.HTMLElement) {
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

				tr.ForEach("tr", func(countTr int, el *colly.HTMLElement) {
					if countTr != 0 {
						iHolder := getHoldersInTable(el)
						if count == 1 {
							holders.TopInstitutionalHolders = append(holders.TopInstitutionalHolders, iHolder)
						} else {
							holders.TopMutualFundHolders = append(holders.TopMutualFundHolders, iHolder)
						}
					}

				})

			}
			// change table
			count++
		})

	})
	c.Visit(url)
	c.Wait()
	return holders
}

func getHoldersInTable(td *colly.HTMLElement) InstitutionalHolders {
	var iHolder InstitutionalHolders
	td.ForEach("td", func(y int, el *colly.HTMLElement) {
		if y == 0 {
			iHolder.Holder = el.Text
		} else if y == 1 {
			iHolder.Shares = uint32(transformToInteger(el.Text))
		} else if y == 2 {
			date, _ := TransformDate(el.Text)
			iHolder.DateReported = date
		} else if y == 3 {
			iHolder.PercentOut = transformToFloat(el.Text)
		} else if y == 4 {
			iHolder.Value = uint64(transformToInteger(el.Text))
		}
	})
	return iHolder
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
	PercentOut   float32   `json:"percentOut"`
	Value        uint64    `json:"value"`
}
