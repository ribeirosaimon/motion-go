package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/scraping"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MotionCache struct {
	Company map[string]*Store
}

type Store struct {
	Info       nosqlDomain.SummaryStock
	Code       string
	expiration time.Time
}

var Cache *MotionCache

func NewMotionCache(db *mongo.Client) *MotionCache {
	ctx := context.Background()
	if Cache == nil {
		Cache = &MotionCache{
			Company: make(map[string]*Store),
		}
		Cache.cron(ctx, db)
		return Cache
	}
	Cache.cron(ctx, db)
	return Cache
}

func (m *MotionCache) Get(i string) *Store {
	return m.Company[i]
}

func (m *MotionCache) Add(companyCode string) {
	if len(m.Company) >= 50 {
		summary := scraping.GetStockSummary(companyCode)
		var store = &Store{
			Info:       summary,
			Code:       summary.CompanyCode,
			expiration: time.Now().Add(time.Minute * 5),
		}
		m.Company[store.Code] = store
	}
}

func (m *MotionCache) cron(ctx context.Context, mongoConnection *mongo.Client) {
	for {
		time.Sleep(time.Second)
		func(s string) {
			companyRepository := repository.NewSummaryStockRepository(ctx, mongoConnection)
			summary := scraping.GetStockSummary(s)
			if !companyRepository.ExistByField("companyCode", s) {
				summary.Id = primitive.NewObjectID()
				summary.CreatedAt = time.Now()
				summary.UpdatedAt = time.Now()
				companyRepository.Save(summary)
			}
		}("meli")
		fmt.Println("tudo certo")
	}
}
