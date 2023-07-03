package scraping

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	conn *db.Connections
}

func NewScrapingService(conn *db.Connections) *Service {
	return &Service{conn: conn}
}

func (s *Service) GetSummaryStock(stock string) nosqlDomain.SummaryStock {

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

	add := time.Now().Add(time.Minute * 2)
	if add.After(foundCompany.UpdatedAt) {
		log.Printf(fmt.Sprintf("\033[0m Scraping:\033[0m Create scraping in stock: %s.\"", stock))
		summary := getStockSummary(stock)
		summary.Id = foundCompany.Id
		summary.UpdatedAt = time.Now()
		foundCompany, _ = companyRepository.Save(summary)
		return foundCompany
	}
	return foundCompany
}

func (s *Service) GetAllStocks() []string {
	companyRepository := repository.NewSummaryStockRepository(context.Background(), s.conn.GetMongoTemplate())
	ctx := context.Background()
	// Definir o filtro e projeção
	filter := bson.M{}
	projection := bson.M{"_id": 0, "companyCode": 1}

	cur, err := companyRepository.GetCollection().Find(ctx, filter, options.Find().SetProjection(projection))
	// Iterar sobre os documentos retornados
	response := make([]string, 0)
	if err != nil {
		return response
	}
	for cur.Next(ctx) {
		var stock stockDTO
		err := cur.Decode(&stock)
		if err != nil {
			log.Fatal(err)
		}

		// Acessar o valor do campo "companyName"
		response = append(response, stock.CompanyCode)
	}
	return response
}

type stockDTO struct {
	CompanyCode string `bson:"companyCode"`
}
