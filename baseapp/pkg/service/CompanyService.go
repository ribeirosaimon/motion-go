package service

import (
	"database/sql"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/scraping"
)

type CompanyService struct {
	summaryStockRepository *repository.MotionNoSQLRepository[nosqlDomain.SummaryStock]
	close                  *sql.DB
}

func NewCompanyService(conn *db.Connections) *CompanyService {
	return &CompanyService{
		summaryStockRepository: repository.NewSummaryStockRepository(conn.Context, conn.GetMongoTemplate()),
	}
}

func (s *CompanyService) GetCompany(id string) (nosqlDomain.SummaryStock, error) {

	byId, err := s.summaryStockRepository.FindById(id)
	if err != nil || byId.Status == domain.INACTIVE {
		return nosqlDomain.SummaryStock{}, err
	}
	return byId, nil
}

func (s *CompanyService) DeleteCompany(id int64) bool {
	product, err := s.summaryStockRepository.FindById(id)
	if err != nil {
		return false
	}
	product.Status = domain.INACTIVE
	_, err = s.summaryStockRepository.Save(product)
	if err != nil {
		return false
	}
	return true
}

func (s *CompanyService) FindByCompanyCode(companyName string) (nosqlDomain.SummaryStock, error) {
	if !scraping.GetTimeOpenMarket() {
		summaryStock, err := s.summaryStockRepository.FindByField("companyCode", companyName)
		if err != nil {
			return middleware.GetCache().Get(companyName), nil
		}
		return summaryStock, nil
	}
	return middleware.GetCache().Get(companyName), nil
}

func (s *CompanyService) FindAllCompany(limit, page uint32) ([]nosqlDomain.SummaryStock, error) {
	allCompany, err := s.summaryStockRepository.FindAll(int(limit), int(page))
	if err != nil {
		return []nosqlDomain.SummaryStock{}, nil
	}
	return allCompany, nil
}
