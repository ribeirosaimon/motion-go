package service

import (
	"context"
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/middleware"

	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type ShoppingCartService struct {
	shoppingCartRepository repository.MotionRepository[nosqlDomain.ShoppingCart]
	profileService         ProfileService
	companyService         CompanyService
}

func NewShoppingCartService(ctx context.Context, c *db.Connections) ShoppingCartService {
	return ShoppingCartService{
		shoppingCartRepository: repository.NewShoppingCartRepository(ctx, c.GetMongoTemplate()),
		profileService:         NewProfileService(c),
		companyService:         NewCompanyService(c),
	}
}

func (s ShoppingCartService) GetShoppingCart(user middleware.LoggedUser) (nosqlDomain.ShoppingCart, error) {
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

func (s ShoppingCartService) CreateShoppingCart(loggedUser middleware.LoggedUser) (nosqlDomain.ShoppingCart, error) {
	_, err := s.GetShoppingCart(loggedUser)
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

func (s ShoppingCartService) DeleteShoppingCart(loggedUser middleware.LoggedUser) error {
	shoppingCart, err := s.GetShoppingCart(loggedUser)
	if err != nil {
		return errors.New("you not have a shooping cart")
	}
	err = s.shoppingCartRepository.DeleteById(shoppingCart.Id)
	if err != nil {
		return errors.New("problem when delete your shopping cart")
	}
	return nil
}

func (s ShoppingCartService) AddProductInShoppingCart(loggedUser middleware.LoggedUser, id int64) (nosqlDomain.ShoppingCart, error) {
	shoppingCart, err := s.GetShoppingCart(loggedUser)
	if err != nil {
		return nosqlDomain.ShoppingCart{}, err
	}
	productDb, err := s.companyService.GetCompany(id)
	if err != nil {
		return nosqlDomain.ShoppingCart{}, err
	}
	shoppingCart.Products = append(shoppingCart.Products, productDb)

	shoppingCart.UpdatedAt = time.Now()

	return s.shoppingCartRepository.Save(shoppingCart)
}
