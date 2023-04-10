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

func (s service) saveProduct(dto ProductDto) (domain.Product, error) {
	var product domain.Product

	product.Name = dto.Name
	product.Image = dto.Image
	product.Price = dto.Price
	product.Status = domain.ACTIVE
	product.CreatedAt = time.Now()

	return s.productRepository.Save(product)
}
