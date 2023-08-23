package service

import (
	"context"
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/dto"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain"
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
	exist := s.portfolioRepository.ExistByField("ownerId", user.UserId)
	if !exist {
		return nosqlDomain.Portfolio{}, errors.New("not found")
	}
	cart, err := s.portfolioRepository.FindByField("ownerId", user.UserId)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}
	return cart, nil
}

func (s PortfolioService) CreatePortfolio(loggedUser middleware.LoggedUser) (nosqlDomain.Portfolio, error) {
	_, err := s.GetPortfolio(loggedUser)
	if err == nil {
		return nosqlDomain.Portfolio{}, errors.New("error in loggedUser")
	}
	if s.portfolioRepository.ExistByField("ownerId", loggedUser.UserId) {
		return nosqlDomain.Portfolio{}, errors.New("you already have a shopping cart")
	}

	var portfolio nosqlDomain.Portfolio

	user, err := s.profileService.FindProfileByUserId(loggedUser.UserId)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}

	portfolio.Id = primitive.NewObjectID()
	portfolio.OwnerId = user.GetId().(uint64)
	portfolio.CreatedAt = time.Now()
	portfolio.Status = domain.ACTIVE
	portfolio.Companies = make([]nosqlDomain.MineStock, 0)
	savedShoppingCart, err := s.portfolioRepository.Save(portfolio)
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

func (s PortfolioService) AddCompanyInPortfolioById(loggedUser middleware.LoggedUser, id string, buyPrice dto.BuyPriceDTO) (nosqlDomain.Portfolio, error) {
	portfolio, err := s.GetPortfolio(loggedUser)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}
	companyDb, err := s.companyService.GetCompany(id)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}
	for _, v := range portfolio.Companies {
		if v.StockId == companyDb.Id {
			return nosqlDomain.Portfolio{}, errors.New("company already exist in your portfolio")
		}
	}
	var mineStock = nosqlDomain.MineStock{
		StockId:  companyDb.Id,
		BuyPrice: buyPrice.Price,
		Quantity: buyPrice.Quantity,
	}
	portfolio.Price += mineStock.CalculeValue()
	portfolio.Companies = append(portfolio.Companies, mineStock)

	portfolio.UpdatedAt = time.Now()

	return s.portfolioRepository.Save(portfolio)
}

func (s PortfolioService) AddCompanyInPortfolioByCode(loggedUser middleware.LoggedUser, companyCode string, buyPrice dto.BuyPriceDTO) (nosqlDomain.Portfolio, error) {
	portfolio, err := s.GetPortfolio(loggedUser)
	if err != nil {
		return nosqlDomain.Portfolio{}, errors.New("you not have a portfolio")
	}
	companyDb, err := s.companyService.FindByCompanyByCodeOrName(companyCode, true)

	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}
	var mineStock = nosqlDomain.MineStock{
		StockId:  companyDb.Id,
		BuyPrice: buyPrice.Price,
		Quantity: buyPrice.Quantity,
	}
	portfolio.Price += mineStock.CalculeValue()
	portfolio.Companies = append(portfolio.Companies, mineStock)

	save, err := s.portfolioRepository.Save(portfolio)
	if err != nil {
		return nosqlDomain.Portfolio{}, err
	}
	return save, nil
}
