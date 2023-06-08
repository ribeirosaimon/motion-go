package scraping

import (
	"fmt"

	"github.com/gocolly/colly"
)

func financials(ticket string) map[string]map[string]interface{} {
	financialsType := []string{"financials", "balance-sheet", "cash-flow"}
	result := make(map[string]map[string]interface{})

	for _, s := range financialsType {
		url := fmt.Sprintf("%s/quote/%s/%s", domain, ticket, s)
		yearsResults := make(map[string]interface{})

		c := prepareColly()

		c.OnHTML("div[class='D(tbr) C($primaryColor)']", func(e *colly.HTMLElement) {
			var count = 0
			e.ForEach("span", func(v int, el *colly.HTMLElement) {

				if v > 1 {
					date, err := TransformDate(el.Text)
					if err == nil {
						yearsResults[fmt.Sprintf("%d", date.Year())] = count
						count++
					}
				}

			})
		})

		c.OnHTML("div[data-test='fin-row']", func(e *colly.HTMLElement) {
			var keyMap string
			breakdown := make(map[string]interface{})
			e.ForEach("span", func(spanCount int, el *colly.HTMLElement) {
				if spanCount == 0 {
					if el.Text == "Diluted EPS" {
						fmt.Sprintf(" ok")
					}
					keyMap = el.Text
				} else {
					if spanCount == 1 {
						breakdown[getKeyByValue(yearsResults, 0)] = transformToInteger(el.Text)
					} else if spanCount == 2 {
						breakdown[getKeyByValue(yearsResults, 1)] = transformToInteger(el.Text)
					} else if spanCount == 3 {
						breakdown[getKeyByValue(yearsResults, 2)] = transformToInteger(el.Text)
					} else if spanCount == 4 {
						breakdown[getKeyByValue(yearsResults, 3)] = transformToInteger(el.Text)
					}
				}
			})

			result[keyMap] = breakdown
		})

		c.Visit(url)
		c.Wait()
	}

	return result
}
func getKeyByValue(m map[string]interface{}, value interface{}) string {
	for key, val := range m {
		if val == value {
			return key
		}
	}
	return ""
}
