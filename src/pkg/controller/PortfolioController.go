package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpResponse"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/src/pkg/dto"
	"github.com/ribeirosaimon/motion-go/src/pkg/service"
)

type portfolioController struct {
	portfolioService *service.PortfolioService
}

func NewPortfolioController() *portfolioController {
	shoppingCartService := service.NewPortfolioService(context.Background(), db.Conn)

	return &portfolioController{portfolioService: &shoppingCartService}
}

func (s *portfolioController) CreatePortfolio(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	_, err = s.portfolioService.CreatePortfolio(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpResponse.Entity(c, http.StatusCreated, nil)
}

func (s *portfolioController) GetPortfolio(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	cart, err := s.portfolioService.GetPortfolio(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpResponse.Entity(c, http.StatusOK, cart)
}

func (s *portfolioController) ExcludePortfolio(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	err = s.portfolioService.DeletePortfolio(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpResponse.Entity(c, http.StatusOK, nil)
}

func (s *portfolioController) AddCompanyInPortfolio(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)

	var price dto.BuyPriceDTO
	if err := c.BindJSON(&price); err != nil {
		exceptions.BodyError().Throw(c)
		return
	}

	if err != nil {
		exceptions.BodyError().Throw(c)
		return
	}

	cart, err := s.portfolioService.AddCompanyInPortfolioById(loggedUser, c.Param("id"), price)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpResponse.Entity(c, http.StatusOK, cart)
}

func (s *portfolioController) AddCompanyByCodeInPortfolio(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	var price dto.BuyPriceDTO
	if err := c.BindJSON(&price); err != nil {
		exceptions.BodyError().Throw(c)
		return
	}
	portfolio, err := s.portfolioService.AddCompanyInPortfolioByCode(loggedUser, c.Param("companyCode"), price)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpResponse.Entity(c, http.StatusOK, portfolio)
}
