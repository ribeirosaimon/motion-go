package product

import (
	"database/sql"
	"time"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/repository"
	"gorm.io/gorm"
)

type service struct {
	productRepository repository.MotionRepository[domain.Product]
	close             *sql.DB
}

func NewShoppingCartService(conn *gorm.DB, close *sql.DB) service {
	return service{
		productRepository: repository.NewProductRepository(conn),
		close:             close,
	}
}

func (s service) getProduct(id string) (domain.Product, error) {
	byId, err := s.productRepository.FindById(id)
	if err != nil || byId.Status == domain.INACTIVE {
		return domain.Product{}, err
	}
	return byId, nil
}

func (s service) saveProduct(dto ProductDto) (domain.Product, error) {
	var product domain.Product

	product.Name = dto.Name
	product.Image = dto.Image
	product.Price = dto.Price
	product.Status = domain.ACTIVE
	product.CreatedAt = time.Now()

	return s.productRepository.Save(product)
}

func (s service) updateProduct(dto ProductDto, id string) (domain.Product, error) {
	product, err := s.getProduct(id)
	if err != nil {
		return domain.Product{}, err
	}

	product.Name = dto.Name
	product.Image = dto.Image
	product.Price = dto.Price
	product.UpdatedAt = time.Now()

	return s.productRepository.Save(product)
}
