package shoppingcart

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/domain/nosql"
	"github.com/ribeirosaimon/motion-go/pkg/product"
	"github.com/ribeirosaimon/motion-go/pkg/profile"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"github.com/ribeirosaimon/motion-go/repository"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type service struct {
	shoppingCartRepository repository.MotionRepository[nosql.ShoppingCart]
	profileService         profile.Service
	productService         product.Service
	close                  *sql.DB
}

func NewShoppingCartService(conn *gorm.DB, close *sql.DB) service {
	return service{
		shoppingCartRepository: repository.NewShoppingCartRepository(conn),
		profileService:         profile.NewProfileService(conn, close),
		productService:         product.NewProductService(conn, close),
		close:                  close,
	}
}

func (s service) getShoppingCart(user security.LoggedUser) (nosql.ShoppingCart, error) {
	exist := s.shoppingCartRepository.ExistByField("profile_id", user.UserId)
	if !exist {
		return nosql.ShoppingCart{}, errors.New("not found")
	}
	cart, err := s.shoppingCartRepository.FindByField("profile_id", user.UserId)
	if err != nil {
		return nosql.ShoppingCart{}, err
	}
	return cart, nil
}

func (s service) createShoppingCart(loggedUser security.LoggedUser) (nosql.ShoppingCart, error) {
	_, err := s.getShoppingCart(loggedUser)
	if err == nil {
		return nosql.ShoppingCart{}, errors.New("you already have a shopping cart")
	}

	var shoppingCart nosql.ShoppingCart

	user, err := s.profileService.FindProfileByUserId(loggedUser.UserId)
	if err != nil {
		return nosql.ShoppingCart{}, err
	}

	shoppingCart.ProfileId = loggedUser.UserId
	shoppingCart.Owner = user
	shoppingCart.CreatedAt = time.Now()
	savedShoppingCart, err := s.shoppingCartRepository.Save(shoppingCart)
	if err != nil {
		return nosql.ShoppingCart{}, err
	}
	return savedShoppingCart, nil
}

func (s service) deleteShoppingCart(loggedUser security.LoggedUser) error {
	shoppingCart, err := s.getShoppingCart(loggedUser)
	if err != nil {
		return errors.New("you not have a shooping cart")
	}
	err = s.shoppingCartRepository.DeleteById(shoppingCart.Id)
	if err != nil {
		return errors.New("problem when delete your shopping cart")
	}
	return nil
}

func (s service) addProductInShoppingCart(loggedUser security.LoggedUser, productDTO productDTO) (nosql.ShoppingCart, error) {
	shoppingCart, err := s.getShoppingCart(loggedUser)
	if err != nil {
		return nosql.ShoppingCart{}, err
	}
	productDb, err := s.productService.GetProduct(productDTO.Id)
	if err != nil {
		return nosql.ShoppingCart{}, err
	}
	shoppingCart.Products = append(shoppingCart.Products, productDb)

	values := productDb.Price.Mul(decimal.NewFromInt(productDTO.Amount))
	shoppingCart.Price = values
	shoppingCart.UpdatedAt = time.Now()

	return s.shoppingCartRepository.Save(shoppingCart)
}
