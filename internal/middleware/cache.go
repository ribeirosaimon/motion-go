package middleware

import (
	"github.com/ribeirosaimon/motion-go/internal/db"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/scraping"
)

type MotionCache struct {
	Company map[string]*Store
	Service *scraping.Service
}

type Store struct {
	Info       nosqlDomain.SummaryStock
	Code       string
	expiration time.Time
}

var Cache *MotionCache

func NewMotionCache(conn *db.Connections) *MotionCache {
	if Cache == nil {
		Cache = &MotionCache{
			Company: make(map[string]*Store),
			Service: scraping.NewScrapingService(conn),
		}
		Cache.cron()
		return Cache
	}
	Cache.cron()
	return Cache
}

func (m *MotionCache) Get(i string) *Store {
	return m.Company[i]
}

func (m *MotionCache) Add(companyCode string) {
	if len(m.Company) >= 50 {
		summary := m.Service.GetSummaryStock(companyCode)
		var store = &Store{
			Info:       summary,
			Code:       summary.CompanyCode,
			expiration: time.Now().Add(time.Minute * 5),
		}
		m.Company[store.Code] = store
	}
}

func (m *MotionCache) cron() {
	for {
		time.Sleep(time.Minute * 15)
		func(s string) {
			m.Service.GetSummaryStock(s)
		}("meli")

	}
}
