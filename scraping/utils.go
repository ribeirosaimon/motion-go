package scraping

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func TransformToPrice(v string) decimal.Decimal {
	if strings.Contains(v, ",") {
		v = strings.ReplaceAll(v, ",", "")

	}
	dec, err := decimal.NewFromString(v)
	if err != nil {
		fmt.Println("Error:", err)
		return decimal.Decimal{}
	}
	return dec
}

func transformToFloat(v string) float32 {
	if strings.Contains(v, "%") {
		v = strings.ReplaceAll(v, "%", "")

	}
	float32Value, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return float32(0)
	}
	return float32(float32Value)
}

func transformToInteger(v string) int {
	if strings.Contains(v, ",") {
		newValue := strings.ReplaceAll(v, ",", "")
		if s, err := strconv.Atoi(newValue); err == nil {
			return s
		}
	}
	return 0
}

func transformDate(layout, dateString string) (time.Time, error) {
	parse, err2 := time.Parse("Jan 3 2006 00:00", "Mar 30 2023")

	fmt.Println(parse, err2)
	date, err := time.Parse(layout, dateString+" 00:00")
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
