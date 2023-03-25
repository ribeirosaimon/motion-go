package shoppingcart

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/profile"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"github.com/ribeirosaimon/motion-go/repository"
	"gorm.io/gorm"
)

type service struct {
	shoppingCartRepository repository.MotionRepository[domain.ShoppingCart]
	profileService         profile.Service
	close                  *sql.DB
}

func NewShoppingCartService(conn *gorm.DB, close *sql.DB) service {
	return service{
		shoppingCartRepository: repository.NewShoppingCartRepository(conn),
		profileService:         profile.NewProfileService(conn, close),
		close:                  close,
	}
}

func (s service) loadShoppingCart(user security.LoggedUser) (domain.ShoppingCart, error) {
	exist := s.shoppingCartRepository.ExistByField("profile_id", user.UserId)
	if !exist {
		return domain.ShoppingCart{}, errors.New("not found")
	}
	cart, err := s.shoppingCartRepository.FindByField("profile_id", user.UserId)
	if err != nil {
		return domain.ShoppingCart{}, err
	}
	return cart, nil
}

func (s service) createShoppingCart(loggedUser security.LoggedUser) (domain.ShoppingCart, error) {
	_, err := s.loadShoppingCart(loggedUser)
	if err == nil {
		return domain.ShoppingCart{}, errors.New("you already have a shopping cart")
	}

	var shoppingCart domain.ShoppingCart

	user, err := s.profileService.FindProfileByUserId(loggedUser.UserId)
	if err != nil {
		return domain.ShoppingCart{}, err
	}

	shoppingCart.ProfileId = loggedUser.UserId
	shoppingCart.Owner = user
	shoppingCart.CreatedAt = time.Now()
	savedShoppingCart, err := s.shoppingCartRepository.Save(shoppingCart)
	if err != nil {
		return domain.ShoppingCart{}, err
	}
	return savedShoppingCart, nil
}
