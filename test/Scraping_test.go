package test

import (
	"testing"

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
