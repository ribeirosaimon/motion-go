package scraping

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func TransformToPrice(v string) decimal.Decimal {

	v = strings.TrimSpace(v)
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
	re := regexp.MustCompile("[0-9.]+")
	matches := re.FindAllString(v, -1)
	result := ""
	for _, match := range matches {
		result += match
	}

	float32Value, err := strconv.ParseFloat(result, 32)
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

func TransformDate(dateString string) (time.Time, error) {
	var err error
	var year, day int
	var month time.Month

	if strings.Contains(dateString, "/") {
		splitedDate := strings.Split(dateString, "/")
		day, err = strconv.Atoi(splitedDate[1])
		monthInt, _ := strconv.Atoi(splitedDate[0])
		month = time.Month(monthInt)
		year, err = strconv.Atoi(splitedDate[2])
	} else if strings.Contains(dateString, "-") {
		dateString = strings.Split(dateString, " ")[0]
		split := strings.Split(dateString, "-")

		day, err = strconv.Atoi(split[2])
		month, err = monthAbbreviationToNumber(split[1])
		year, err = strconv.Atoi(split[0])

	} else {
		replacedString := strings.ReplaceAll(dateString, ",", "")
		split := strings.Split(replacedString, " ")

		day, err = strconv.Atoi(split[1])
		month, err = monthAbbreviationToNumber(split[0])
		year, err = strconv.Atoi(split[2])
	}
	if err != nil {
		return time.Time{}, err
	}
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	return date, nil
}

func monthAbbreviationToNumber(monthAbbreviation string) (time.Month, error) {
	t, err := time.Parse("Jan", monthAbbreviation)
	if err != nil {
		return 0, err
	}
	return t.Month(), nil
}
