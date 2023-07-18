package middleware

import (
	"log"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"

	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/scraping"
)

type MotionCache struct {
	Company   map[string]*Store
	service   *scraping.Service
	CacheTime uint8
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
			Company:   make(map[string]*Store),
			service:   service,
			CacheTime: cacheTime,
		}
		Cache.cron(haveScraping, scrapingTime)
		return Cache
	}
	Cache.cron(haveScraping, scrapingTime)
	return Cache
}

func (m *MotionCache) Get(i string) nosqlDomain.SummaryStock {
	if m.contains(i) {
		return m.Company[i].Info
	}
	summaryStock := m.service.GetSummaryStock(i)
	m.Add(summaryStock)
	return summaryStock
}

func (m *MotionCache) Add(company nosqlDomain.SummaryStock) {
	if len(m.Company) <= 50 {
		var store = &Store{
			Info:       company,
			Code:       company.CompanyCode,
			expiration: time.Now().Add(time.Minute * time.Duration(m.CacheTime)),
		}
		m.Company[store.Code] = store
	}
}

func (m *MotionCache) cron(haveScraping bool, scrapingTime uint8) {

	if scraping.GetTimeOpenMarket() {
		if haveScraping {
			for {
				stocks := m.service.GetAllStocks()
				for _, stock := range stocks {
					cacheCompany := m.Company[stock]
					if cacheCompany == nil || cacheCompany.expiration.Before(time.Now()) {
						func(s string) {
							summaryStock := m.service.GetSummaryStock(s)
							m.Add(summaryStock)
						}(stock)
					}
				}
				log.Println("cron was finished")
				time.Sleep(time.Minute * time.Duration(scrapingTime))
			}
		}
	} else {
		log.Println("close market")
	}

}

func (m *MotionCache) contains(company string) bool {
	for _, v := range m.Company {
		if v.Code == company {
			return true
		}
	}
	return false
}
