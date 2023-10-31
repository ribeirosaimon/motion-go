package middleware

import (
	"errors"
	"log"
	"time"

	"github.com/ribeirosaimon/motion-go/api/internal/config"
	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/src/scraping"
	"github.com/ribeirosaimon/motion-go/confighub/domain/nosqlDomain"
)

type MotionCache struct {
	Service    *scraping.Service
	Config     *config.MotionConfig
	NextModify *int64
	Company    map[string]*Store
}

type Store struct {
	Info       nosqlDomain.SummaryStock
	expiration time.Time
}

var Cache *MotionCache

func GetCache() *MotionCache {
	return Cache
}

func NewMotionCache(conn *db.Connections) *MotionCache {
	motionConfig := config.GetMotionConfig()
	if Cache == nil {
		service := scraping.NewScrapingService(conn)
		Cache = &MotionCache{
			Company:    make(map[string]*Store),
			Service:    service,
			NextModify: nil,
			Config:     &motionConfig,
		}
		Cache.cron()
		return Cache
	}
	Cache.cron()
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
		summaryStock, err := m.Service.GetSummaryStock(i)
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
			expiration: time.Now().Add(time.Minute * time.Duration(m.Config.CacheTime)),
		}
		m.Company[store.Info.CompanyCode] = store
	}
}

func (m *MotionCache) cron() {

	if scraping.GetTimeOpenMarket() {

		if m.Config.HaveScraping {
			for {
				stocks := m.Service.GetAllStocks()
				for _, stock := range stocks {
					cacheCompany := m.Company[stock]
					if cacheCompany == nil || cacheCompany.expiration.Before(time.Now()) {
						func(s string) {
							getSummaryStock, err := m.Service.GetSummaryStock(s)
							if err == nil {
								unix := time.Now().Unix()
								unix += int64(m.Config.ScrapingTime * 60)
								summaryStock := getSummaryStock
								m.Add(summaryStock)
								m.NextModify = &unix
							}
						}(stock)
					}
				}
				log.Println("cron was finished")
				time.Sleep(time.Minute * time.Duration(m.Config.ScrapingTime))
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
