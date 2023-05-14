package shoppingcart

import (
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/product"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/profile"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/internal/security"
	"github.com/shopspring/decimal"
)

type service struct {
	shoppingCartRepository repository.MotionRepository[nosqlDomain.ShoppingCart]
	profileService         profile.Service
	productService         product.Service
}

func NewShoppingCartService(c *db.Connections) service {
	return service{
		shoppingCartRepository: repository.NewShoppingCartRepository(c.GetMongoTemplate()),
		profileService:         profile.NewProfileService(c),
		productService:         product.NewProductService(c),
	}
}

func (s service) getShoppingCart(user security.LoggedUser) (nosqlDomain.ShoppingCart, error) {
	exist := s.shoppingCartRepository.ExistByField("profile_id", user.UserId)
	if !exist {
		return nosqlDomain.ShoppingCart{}, errors.New("not found")
	}
	cart, err := s.shoppingCartRepository.FindByField("profile_id", user.UserId)
	if err != nil {
		return nosqlDomain.ShoppingCart{}, err
	}
	return cart, nil
}

func (s service) createShoppingCart(loggedUser security.LoggedUser) (nosqlDomain.ShoppingCart, error) {
	_, err := s.getShoppingCart(loggedUser)
	if err == nil {
		return nosqlDomain.ShoppingCart{}, errors.New("you already have a shopping cart")
	}

	var shoppingCart nosqlDomain.ShoppingCart

	user, err := s.profileService.FindProfileByUserId(loggedUser.UserId)
	if err != nil {
		return nosqlDomain.ShoppingCart{}, err
	}

	shoppingCart.ProfileId = loggedUser.UserId
	shoppingCart.Owner = user
	shoppingCart.CreatedAt = time.Now()
	savedShoppingCart, err := s.shoppingCartRepository.Save(shoppingCart)
	if err != nil {
		return nosqlDomain.ShoppingCart{}, err
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

func (s service) addProductInShoppingCart(loggedUser security.LoggedUser, productDTO productDTO) (nosqlDomain.ShoppingCart, error) {
	shoppingCart, err := s.getShoppingCart(loggedUser)
	if err != nil {
		return nosqlDomain.ShoppingCart{}, err
	}
	productDb, err := s.productService.GetProduct(productDTO.Id)
	if err != nil {
		return nosqlDomain.ShoppingCart{}, err
	}
	shoppingCart.Products = append(shoppingCart.Products, productDb)

	values := productDb.Price.Mul(decimal.NewFromInt(productDTO.Amount))
	shoppingCart.Price = values
	shoppingCart.UpdatedAt = time.Now()

	return s.shoppingCartRepository.Save(shoppingCart)
}
