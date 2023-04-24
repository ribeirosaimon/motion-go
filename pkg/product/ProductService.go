package product

import (
	"database/sql"
	"errors"
	"time"

	sql2 "github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/repository"
	"gorm.io/gorm"
)

type Service struct {
	productRepository repository.MotionRepository[sql2.Product]
	close             *sql.DB
}

func NewProductService(conn *gorm.DB, close *sql.DB) Service {
	return Service{
		productRepository: repository.NewProductRepository(conn),
		close:             close,
	}
}

func (s Service) GetProduct(id int64) (sql2.Product, error) {
	byId, err := s.productRepository.FindById(id)
	if err != nil || byId.Status == sql2.INACTIVE {
		return sql2.Product{}, err
	}
	return byId, nil
}

func (s Service) saveProduct(dto ProductDto) (sql2.Product, error) {
	var product sql2.Product

	if s.productRepository.ExistByField("name", dto.Name) {
		return product, errors.New("this product already exists")
	}

	product.Name = dto.Name
	product.Image = dto.Image
	product.Price = dto.Price
	product.Status = sql2.ACTIVE
	product.CreatedAt = time.Now()

	return s.productRepository.Save(product)
}

func (s Service) updateProduct(dto ProductDto, id int64) (sql2.Product, error) {
	product, err := s.GetProduct(id)
	if err != nil {
		return sql2.Product{}, errors.New("product not found")
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
	product.Status = sql2.INACTIVE
	_, err = s.productRepository.Save(product)
	if err != nil {
		return false
	}
	return true
}
