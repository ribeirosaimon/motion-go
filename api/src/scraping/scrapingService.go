package scraping

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/grpcconnection"
	"github.com/ribeirosaimon/motion-go/api/internal/repository"
	"github.com/ribeirosaimon/motion-go/confighub/domain/nosqlDomain"
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

func (s *Service) GetSummaryStock(stock string) (nosqlDomain.SummaryStock, error) {

	companyRepository := repository.NewSummaryStockRepository(context.Background(), s.conn.GetMongoTemplate())
	foundCompany, err := companyRepository.FindByField("companyCode", stock)

	if err != nil {
		mongoSummaryStock, err := GetStockSummary(stock)
		if err != nil {
			return nosqlDomain.SummaryStock{}, err
		}

		foundCompany, _ = companyRepository.Save(mongoSummaryStock)
		return foundCompany, nil
	}

	add := time.Now().Add(time.Minute * 2)
	if add.After(foundCompany.UpdatedAt) {
		log.Printf(fmt.Sprintf("\033[0m Scraping:\033[0m Create scraping in stock: %s.\"", stock))
		mongoSummaryStock, err := GetStockSummary(stock)
		if err != nil {
			return nosqlDomain.SummaryStock{}, err
		}

		foundCompany, _ = companyRepository.Save(mongoSummaryStock)
		return foundCompany, nil
	}
	return foundCompany, nil
}

func GetStockSummary(stock string) (nosqlDomain.SummaryStock, error) {
	newStock, err := grpcconnection.GetStock(stock, false)
	if err != nil {
		return nosqlDomain.SummaryStock{}, err
	}

	if newStock.StockValue.GetPrice() == float64(0) {
		return nosqlDomain.SummaryStock{}, errors.New("this stock does not exist")
	}

	summary := newStock
	summary.Id = primitive.NewObjectID().Hex()
	mongoSummaryStock := nosqlDomain.ChangeProtoToMongo(summary)
	return *mongoSummaryStock, nil
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
	IsBrazilian bool   `bson:"isBrazilian"`
}
