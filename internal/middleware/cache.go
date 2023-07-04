package middleware

import (
	"log"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"

	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/scraping"
)

type MotionCache struct {
	Company map[string]*Store
	service *scraping.Service
}

type Store struct {
	Info       nosqlDomain.SummaryStock
	Code       string
	expiration time.Time
}

var Cache *MotionCache

func GetCache() *MotionCache {
	return Cache
}

func NewMotionCache(conn *db.Connections, haveScraping bool, scrapingTime, cacheTime uint8) *MotionCache {
	if Cache == nil {
		service := scraping.NewScrapingService(conn)
		Cache = &MotionCache{
			Company: make(map[string]*Store),
			service: service,
		}
		Cache.cron(haveScraping, scrapingTime, cacheTime)
		return Cache
	}
	Cache.cron(haveScraping, scrapingTime, cacheTime)
	return Cache
}

func (m *MotionCache) Get(i string) *Store {
	return m.Company[i]
}

func (m *MotionCache) Add(cacheTime uint8, company nosqlDomain.SummaryStock) {
	if len(m.Company) <= 50 {
		var store = &Store{
			Info:       company,
			Code:       company.CompanyCode,
			expiration: time.Now().Add(time.Minute * time.Duration(cacheTime)),
		}
		m.Company[store.Code] = store
	}
}

func (m *MotionCache) cron(haveScraping bool, scrapingTime, cacheTime uint8) {
	if haveScraping {
		for {
			stocks := m.service.GetAllStocks()
			for _, stock := range stocks {
				cacheCompany := Cache.Get(stock)
				if cacheCompany == nil || cacheCompany.expiration.Before(time.Now()) {
					func(s string) {
						summaryStock := m.service.GetSummaryStock(s)
						m.Add(cacheTime, summaryStock)
					}(stock)
				}
			}
			log.Println("CRON FINISHED")
			time.Sleep(time.Minute * time.Duration(scrapingTime))
		}
	}
}
