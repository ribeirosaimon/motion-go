package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/service"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
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
	c.Status(http.StatusCreated)
}

func (s *portfolioController) GetPortfolio(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	cart, err := s.portfolioService.GetPortfolio(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Created(c, cart)
}

func (s *portfolioController) ExcludePortfolio(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	err = s.portfolioService.DeletePortfolio(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Ok(c, nil)
}

func (s *portfolioController) AddCompanyInPortfolio(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(c)
		return
	}

	cart, err := s.portfolioService.AddCompanyInPortfolioById(loggedUser, id)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Ok(c, cart)
}

func (s *portfolioController) AddCompanyByCodeInPortfolio(ctx *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(ctx)
	err = s.portfolioService.AddCompanyInPortfolioByCode(loggedUser, ctx.Param("companyName"))
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	httpresponse.Ok(ctx, companyName)
}
