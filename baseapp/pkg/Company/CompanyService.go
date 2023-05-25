package Company

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"

	"github.com/ribeirosaimon/motion-go/internal/domain"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type Service struct {
	companyRepository repository.MotionRepository[sqlDomain.Company]
	close             *sql.DB
}

func NewCompanyService(conn *db.Connections) Service {
	return Service{
		companyRepository: repository.NewCompanyRepository(conn.GetPgsqTemplate()),
	}
}

func (s Service) GetCompany(id int64) (sqlDomain.Company, error) {
	byId, err := s.companyRepository.FindById(id)
	if err != nil || byId.Status == domain.INACTIVE {
		return sqlDomain.Company{}, err
	}
	return byId, nil
}

func (s Service) saveCompany(dto companyDto) (sqlDomain.Company, error) {
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

func (s Service) updateCompany(dto companyDto, id int64) (sqlDomain.Company, error) {
	company, err := s.GetCompany(id)
	if err != nil {
		return sqlDomain.Company{}, errors.New("company not found")
	}

	company.Name = dto.Name
	company.Image = dto.Image
	company.UpdatedAt = time.Now()

	return s.companyRepository.Save(company)
}

func (s Service) deleteCompany(id int64) bool {
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
