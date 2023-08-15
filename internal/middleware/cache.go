package middleware

import (
	"errors"
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

func (m *MotionCache) GetByCompanyName(i string) (nosqlDomain.SummaryStock, error) {
	return m.getCompanyInCache(i, false)
}

func (m *MotionCache) GetByCompanyCode(i string) (nosqlDomain.SummaryStock, error) {
	return m.getCompanyInCache(i, false)
}

func (m *MotionCache) getCompanyInCache(i string, companyCode bool) (nosqlDomain.SummaryStock, error) {
	contains, err := m.contains(i, companyCode)
	if err != nil {
		summaryStock, err := m.service.GetSummaryStock(i)
		if err != nil {
			return nosqlDomain.SummaryStock{}, err
		}
		m.Add(summaryStock)
		return summaryStock, nil
	}
	return contains, nil
}

func (m *MotionCache) Add(company nosqlDomain.SummaryStock) {
	if len(m.Company) <= 50 {
		var store = &Store{
			Info:       company,
			expiration: time.Now().Add(time.Minute * time.Duration(m.CacheTime)),
		}
		m.Company[store.Info.CompanyCode] = store
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
							getSummaryStock, err := m.service.GetSummaryStock(s)
							if err == nil {
								summaryStock := getSummaryStock
								m.Add(summaryStock)
							}
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

func (m *MotionCache) contains(company string, companyCode bool) (nosqlDomain.SummaryStock, error) {
	for _, v := range m.Company {
		if companyCode && (v.Info.CompanyCode == company) {
			return v.Info, nil
		}
		if !companyCode && (v.Info.CompanyName == company) {
			return v.Info, nil
		}
	}
	return nosqlDomain.SummaryStock{}, errors.New("")
}
