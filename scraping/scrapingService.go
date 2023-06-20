package scraping

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	conn *db.Connections
}

func NewScrapingService(conn *db.Connections) *Service {
	return &Service{conn: conn}
}

func (s *Service) GetSummaryStock(stock string) nosqlDomain.SummaryStock {
	log.Printf(fmt.Sprintf("\033[Scraping:\033[0m Create scraping in stock: %s.\"", stock))
	companyRepository := repository.NewSummaryStockRepository(context.Background(), s.conn.GetMongoTemplate())
	foundCompany, err := companyRepository.FindByField("companyCode", stock)

	if err != nil {
		summary := getStockSummary(stock)
		summary.Id = primitive.NewObjectID()
		summary.CreatedAt = time.Now()
		summary.UpdatedAt = time.Now()
		foundCompany, _ = companyRepository.Save(summary)
		return foundCompany
	}

	if foundCompany.UpdatedAt.After(time.Now().Add(time.Hour)) {
		summary := getStockSummary(stock)
		summary.Id = foundCompany.Id
		summary.UpdatedAt = time.Now()
		foundCompany, _ = companyRepository.Save(summary)
		return foundCompany
	}
	return foundCompany
}
