package scraping

import (
	"context"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service struct {
	conn *db.Connections
}

var scServ Service

func NewScrapingService(conn *db.Connections) *Service {
	if &scServ == nil {
		return &Service{conn: conn}
	}
	return &scServ
}

func (s *Service) GetSummaryStock(stock string) nosqlDomain.SummaryStock {
	companyRepository := repository.NewSummaryStockRepository(context.Background(), s.conn.GetMongoTemplate())
	summary := getStockSummary(stock)
	if companyRepository.ExistByField("companyCode", stock) {
		foundCompany, _ := companyRepository.FindByField("companyCode", stock)
		return foundCompany
	} else {
		summary.Id = primitive.NewObjectID()
		summary.CreatedAt = time.Now()
		summary.UpdatedAt = time.Now()
		company, _ := companyRepository.Save(summary)
		return company
	}
}
