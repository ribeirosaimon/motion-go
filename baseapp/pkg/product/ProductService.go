package product

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
	productRepository repository.MotionRepository[sqlDomain.Product]
	close             *sql.DB
}

func NewProductService(conections *db.Connections) Service {
	return Service{
		productRepository: repository.NewProductRepository(conections.GetPgsqTemplate()),
	}
}

func (s Service) GetProduct(id int64) (sqlDomain.Product, error) {
	byId, err := s.productRepository.FindById(id)
	if err != nil || byId.Status == domain.INACTIVE {
		return sqlDomain.Product{}, err
	}
	return byId, nil
}

func (s Service) saveProduct(dto ProductDto) (sqlDomain.Product, error) {
	var product sqlDomain.Product

	if s.productRepository.ExistByField("name", dto.Name) {
		return product, errors.New("this product already exists")
	}

	product.Name = dto.Name
	product.Image = dto.Image
	product.Price = dto.Price
	product.Status = domain.ACTIVE
	product.CreatedAt = time.Now()

	return s.productRepository.Save(product)
}

func (s Service) updateProduct(dto ProductDto, id int64) (sqlDomain.Product, error) {
	product, err := s.GetProduct(id)
	if err != nil {
		return sqlDomain.Product{}, errors.New("product not found")
	}

	product.Name = dto.Name
	product.Image = dto.Image
	product.Price = dto.Price
	product.UpdatedAt = time.Now()

	return s.productRepository.Save(product)
}

func (s Service) deleteProduct(id int64) bool {
	product, err := s.productRepository.FindById(id)
	if err != nil {
		return false
	}
	product.Status = domain.INACTIVE
	_, err = s.productRepository.Save(product)
	if err != nil {
		return false
	}
	return true
}
