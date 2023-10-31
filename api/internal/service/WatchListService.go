package service

import (
	"context"
	"errors"
	"time"

	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/middleware"
	"github.com/ribeirosaimon/motion-go/api/internal/repository"
	"github.com/ribeirosaimon/motion-go/confighub/domain/nosqlDomain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WatchListService struct {
	watchListRepository *repository.MotionNoSQLRepository[nosqlDomain.WatchList]
	profileService      *ProfileService
	companyService      *CompanyService
}

func NewWatchListService(ctx context.Context, c *db.Connections) WatchListService {
	return WatchListService{
		watchListRepository: repository.NewWatchListRepository(ctx, c.GetMongoTemplate()),
		profileService:      NewProfileService(c),
		companyService:      NewCompanyService(c),
	}
}

func (s WatchListService) GetWatchList(user middleware.LoggedUser) (nosqlDomain.WatchList, error) {
	exist := s.watchListRepository.ExistByField("ownerId", user.ProfileId)
	if !exist {
		return nosqlDomain.WatchList{}, errors.New("not found")
	}
	cart, err := s.watchListRepository.FindByField("ownerId", user.ProfileId)
	if err != nil {
		return nosqlDomain.WatchList{}, err
	}
	return cart, nil
}

func (s WatchListService) CreateWatchList(loggedUser middleware.LoggedUser) (nosqlDomain.WatchList, error) {
	_, err := s.GetWatchList(loggedUser)
	if err == nil {
		return nosqlDomain.WatchList{}, errors.New("error in loggedUser")
	}
	if s.watchListRepository.ExistByField("ownerId", loggedUser.ProfileId) {
		return nosqlDomain.WatchList{}, errors.New("you already have a shopping cart")
	}

	var portfolio nosqlDomain.WatchList

	user, err := s.profileService.FindProfileByUserId(loggedUser.ProfileId)
	if err != nil {
		return nosqlDomain.WatchList{}, err
	}

	portfolio.Id = primitive.NewObjectID()
	portfolio.OwnerId = user.GetId().(uint64)
	portfolio.CreatedAt = time.Now()
	portfolio.Companies = make([]primitive.ObjectID, 0)
	savedShoppingCart, err := s.watchListRepository.Save(portfolio)
	if err != nil {
		return nosqlDomain.WatchList{}, err
	}
	return savedShoppingCart, nil
}

func (s WatchListService) DeleteWatchList(loggedUser middleware.LoggedUser) error {
	shoppingCart, err := s.GetWatchList(loggedUser)
	if err != nil {
		return errors.New("you not have a shooping cart")
	}
	err = s.watchListRepository.DeleteById(shoppingCart.Id)
	if err != nil {
		return errors.New("problem when delete your shopping cart")
	}
	return nil
}

func (s WatchListService) AddCompanyInWatchListById(loggedUser middleware.LoggedUser, id string) (nosqlDomain.WatchList, error) {
	portfolio, err := s.GetWatchList(loggedUser)
	if err != nil {
		return nosqlDomain.WatchList{}, err
	}
	companyDb, err := s.companyService.GetCompany(id)
	if err != nil {
		return nosqlDomain.WatchList{}, err
	}
	for _, v := range portfolio.Companies {
		if v == companyDb.Id {
			return nosqlDomain.WatchList{}, errors.New("company already exist in your portfolio")
		}
	}

	portfolio.Companies = append(portfolio.Companies, companyDb.Id)

	portfolio.UpdatedAt = time.Now()

	return s.watchListRepository.Save(portfolio)
}

func (s WatchListService) AddCompanyInWatchListByCode(loggedUser middleware.LoggedUser, companyCode string) (nosqlDomain.WatchList, error) {
	portfolio, err := s.GetWatchList(loggedUser)
	if err != nil {
		return nosqlDomain.WatchList{}, errors.New("you not have a portfolio")
	}
	companyDb, err := s.companyService.FindByCompanyByCodeOrName(companyCode, true)

	if err != nil {
		return nosqlDomain.WatchList{}, err
	}

	portfolio.Companies = append(portfolio.Companies, companyDb.Id)

	save, err := s.watchListRepository.Save(portfolio)
	if err != nil {
		return nosqlDomain.WatchList{}, err
	}
	return save, nil
}
