package service

import (
	"database/sql"
	"time"

	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/middleware"
	"github.com/ribeirosaimon/motion-go/api/internal/repository"
	"github.com/ribeirosaimon/motion-go/api/src/scraping"
	"github.com/ribeirosaimon/motion-go/confighub/domain/nosqlDomain"
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
	if err != nil || byId.Status == nosqlDomain.INACTIVE {
		return nosqlDomain.SummaryStock{}, err
	}
	return byId, nil
}

func (s *CompanyService) DeleteCompany(id string) bool {
	product, err := s.summaryStockRepository.FindById(id)
	if err != nil {
		return false
	}

	product.BasicNoSQL = nosqlDomain.BasicNoSQL{
		UpdatedAt: time.Now(),
	}

	_, err = s.summaryStockRepository.Save(product)
	if err != nil {
		return false
	}
	return true
}

func (s *CompanyService) FindByCompanyByCodeOrName(companyName string, code bool) (nosqlDomain.SummaryStock, error) {
	var foundField string
	if code {
		foundField = "companyCode"
	} else {
		foundField = "companyName"
	}
	if !scraping.GetTimeOpenMarket() {
		// looking for this stock in repository
		summaryStock, err := s.summaryStockRepository.FindByField(foundField, companyName)
		if err != nil {
			if code {
				return middleware.GetCache().GetByCompanyCode(companyName)
			}
			return middleware.GetCache().GetByCompanyName(companyName)
		}
		return summaryStock, nil
	}
	name := nosqlDomain.SummaryStock{}
	var err error
	if code {
		name, err = middleware.GetCache().GetByCompanyCode(companyName)
	} else {
		name, err = middleware.GetCache().GetByCompanyName(companyName)
	}
	if err != nil {
		summaryStock, err := s.summaryStockRepository.FindByField(foundField, companyName)
		if err != nil {
			return middleware.GetCache().GetByCompanyName(companyName)
		}
		return summaryStock, nil
	}
	return name, nil
}

func (s *CompanyService) FindAllCompany(limit, page uint32) ([]nosqlDomain.SummaryStock, error) {
	allCompany, err := s.summaryStockRepository.FindAll(int(limit), int(page))
	if err != nil {
		return []nosqlDomain.SummaryStock{}, nil
	}
	return allCompany, nil
}
