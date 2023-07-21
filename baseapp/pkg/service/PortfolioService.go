package service

import (
	"context"
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type PortfolioService struct {
	portfolioRepository *repository.MotionNoSQLRepository[nosqlDomain.Portfolio]
	profileService      *ProfileService
	companyService      *CompanyService
}

func NewPortfolioService(ctx context.Context, c *db.Connections) PortfolioService {
	return PortfolioService{
		portfolioRepository: repository.NewPortfolioRepository(ctx, c.GetMongoTemplate()),
		profileService:      NewProfileService(c),
		companyService:      NewCompanyService(c),
	}
}

func (s PortfolioService) GetPortfolio(user middleware.LoggedUser) (nosqlDomain.Portfolio, error) {
	exist := s.portfolioRepository.ExistByField("owner.id", user.UserId)
	if !exist {
		return nosqlDomain.Portfolio{}, errors.New("not found")
	}
	cart, err := s.portfolioRepository.FindByField("owner.id", user.UserId)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}
	return cart, nil
}

func (s PortfolioService) CreatePortfolio(loggedUser middleware.LoggedUser) (nosqlDomain.Portfolio, error) {
	_, err := s.GetPortfolio(loggedUser)
	if err == nil {
		return nosqlDomain.Portfolio{}, errors.New("you already have a shopping cart")
	}

	var shoppingCart nosqlDomain.Portfolio

	user, err := s.profileService.FindProfileByUserId(loggedUser.UserId)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}

	shoppingCart.Id = primitive.NewObjectID()
	shoppingCart.ProfileId = loggedUser.UserId
	shoppingCart.Owner = user
	shoppingCart.CreatedAt = time.Now()
	savedShoppingCart, err := s.portfolioRepository.Save(shoppingCart)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}
	return savedShoppingCart, nil
}

func (s PortfolioService) DeletePortfolio(loggedUser middleware.LoggedUser) error {
	shoppingCart, err := s.GetPortfolio(loggedUser)
	if err != nil {
		return errors.New("you not have a shooping cart")
	}
	err = s.portfolioRepository.DeleteById(shoppingCart.Id)
	if err != nil {
		return errors.New("problem when delete your shopping cart")
	}
	return nil
}

func (s PortfolioService) AddCompanyInPortfolioById(loggedUser middleware.LoggedUser, id int64) (nosqlDomain.Portfolio, error) {
	portfolio, err := s.GetPortfolio(loggedUser)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}
	companyDb, err := s.companyService.GetCompany(id)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}
	for _, v := range portfolio.Companies {
		if v == companyDb.Id {
			return nosqlDomain.Portfolio{}, errors.New("company already exist in your portfolio")
		}
	}
	portfolio.Companies = append(portfolio.Companies, companyDb.Id)

	portfolio.UpdatedAt = time.Now()

	return s.portfolioRepository.Save(portfolio)
}

func (s PortfolioService) AddCompanyInPortfolioByCode(loggedUser middleware.LoggedUser, companyCode string) error {
	portfolio, err := s.GetPortfolio(loggedUser)
	if err != nil {
		return err
	}
	companyDb, err := s.companyService.FindByCompanyCode(companyCode)

	portfolio.Companies = append(portfolio.Companies, companyDb.Id)
}
