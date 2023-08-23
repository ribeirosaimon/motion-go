package scraping

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScrapingReturnError(t *testing.T) {
	summary := getStockSummary("test")
	assert.Equal(t, summary.Summary.Open, float64(0))
}
