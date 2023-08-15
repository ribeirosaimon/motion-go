package scraping

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScrapingReturnError(t *testing.T) {
	summary, err := getStockSummary("test")
	assert.Equal(t, summary.Summary.Open, float64(0))
	assert.Error(t, err, "this stock does not exist")
}
