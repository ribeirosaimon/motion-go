package scraping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrapingReturnError(t *testing.T) {
	summary := GetStockSummary("test")
	assert.Equal(t, summary.Summary.Open, float64(0))
}
