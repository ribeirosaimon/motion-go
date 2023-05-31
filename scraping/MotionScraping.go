package scraping

import (
	"fmt"
	"github.com/gocolly/colly"
)

func GetStock(stock string) {
	createBasicScrap(stock)
}

func createBasicScrap(v string) {
	domain := fmt.Sprintf("https://finance.yahoo.com")
	url := fmt.Sprintf("%s/quote/%s", domain, v)

	c := colly.NewCollector(colly.AllowedDomains(url))
	selector := `div[data-symbol="MELI"]`

	var symbol, value string

	c.OnHTML(selector, func(e *colly.HTMLElement) {
		// Extraia o valor do atributo data-symbol
		symbol = e.Attr("data-symbol")

		// Extraia o valor do atributo data-value
		value = e.Attr("data-value")

		// Faça algo com os valores obtidos
		fmt.Println("O valor de MELI é:", value, symbol)
	})

}
