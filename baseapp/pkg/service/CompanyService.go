package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/dto"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/scraping"
)

type CompanyService struct {
	companyRepository repository.MotionRepository[sqlDomain.Company]
	summaryStock repository.MotionNoSQLRepository[nosqlDomain.SummaryStock]
	close             *sql.DB
}

func NewCompanyService(conn *db.Connections) CompanyService {
	return CompanyService{
		companyRepository: repository.NewCompanyRepository(conn.GetPgsqTemplate()),
		summaryStock: repository.NewSummaryStockRepository(conn.Context, conn.GetMongoTemplate()),
	}
}

func (s CompanyService) GetCompany(id int64) (sqlDomain.Company, error) {

	byId, err := s.companyRepository.FindById(id)
	if err != nil || byId.Status == domain.INACTIVE {
		return sqlDomain.Company{}, err
	}
	return byId, nil
}

func (s CompanyService) SaveCompany(dto dto.CompanyDTO) (sqlDomain.Company, error) {
	var company sqlDomain.Company

	if s.companyRepository.ExistByField("name", dto.Name) {
		return company, errors.New("this Company already exists")
	}

	company.Name = dto.Name
	company.Image = dto.Image
	company.Status = domain.ACTIVE
	company.CreatedAt = time.Now()

	return s.companyRepository.Save(company)
}

func (s CompanyService) UpdateCompany(dto dto.CompanyDTO, id int64) (sqlDomain.Company, error) {
	company, err := s.GetCompany(id)
	if err != nil {
		return sqlDomain.Company{}, errors.New("company not found")
	}

	company.Name = dto.Name
	company.Image = dto.Image
	company.UpdatedAt = time.Now()

	return s.companyRepository.Save(company)
}

func (s CompanyService) DeleteCompany(id int64) bool {
	product, err := s.companyRepository.FindById(id)
	if err != nil {
		return false
	}
	product.Status = domain.INACTIVE
	_, err = s.companyRepository.Save(product)
	if err != nil {
		return false
	}
	return true
}

func (s CompanyService) FindByCompanyName(companyName string) nosqlDomain.SummaryStock {
	if scraping.GetTimeOpenMarket() {
		field, err := s.companyRepository.FindByField("companyCode", companyName)
		return field.
	}
	return middleware.GetCache().Get(companyName)
}
