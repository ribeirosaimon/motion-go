package shoppingcart

import (
	"database/sql"
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

func (s service) createShoppingCart(loggedUser security.LoggedUser) (domain.ShoppingCart, error) {
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
