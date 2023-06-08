package test

import (
	"testing"
	"time"

	"github.com/ribeirosaimon/motion-go/scraping"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestTransformToPriceWithDot(t *testing.T) {
	v := "10.0"
	value := scraping.TransformToPrice(v)
	fromString, _ := decimal.NewFromString(v)
	assert.Equal(t, value, fromString)
}

func TestTransformToPriceWithDotAndFloat(t *testing.T) {
	v := "10.259"
	value := scraping.TransformToPrice(v)
	fromString, _ := decimal.NewFromString(v)
	assert.Equal(t, value, fromString)
}

func TestTransformToPriceWithComa(t *testing.T) {
	v := "10,0"
	value := scraping.TransformToPrice(v)
	fromString, _ := decimal.NewFromString("100")
	assert.Equal(t, value, fromString)
}

func TestTransformToPriceWithComaAndFloat(t *testing.T) {
	v := "10,269"
	value := scraping.TransformToPrice(v)
	fromString, _ := decimal.NewFromString("10269")
	assert.Equal(t, value, fromString)
}

func TestTransformToPriceWithComaAndDotWithFloat(t *testing.T) {
	v := "10,000.20"
	value := scraping.TransformToPrice(v)
	fromString, _ := decimal.NewFromString("10000.20")
	assert.Equal(t, value, fromString)
}

func TestTransformToPriceWithError(t *testing.T) {
	v := "10-0"
	value := scraping.TransformToPrice(v)
	fromString, _ := decimal.NewFromString(v)
	assert.Equal(t, value, fromString)
}

func TestTransformDate(t *testing.T) {
	historyDataLayout := "Jan 3, 2006"
	date, _ := scraping.TransformDate(historyDataLayout)
	assert.Equal(t, date.Year(), 2006)
	assert.Equal(t, date.Month(), time.Month(1))
	assert.Equal(t, date.Day(), 3)
}

func TestTransformDateOtherLayout(t *testing.T) {
	historyDataLayout := "01/13/2006"
	date, _ := scraping.TransformDate(historyDataLayout)
	assert.Equal(t, date.Year(), 2006)
	assert.Equal(t, date.Day(), 13)
	assert.Equal(t, date.Month(), time.Month(1))
}
