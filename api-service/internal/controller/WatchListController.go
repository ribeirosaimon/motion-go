package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/response"
	"github.com/ribeirosaimon/motion-go/internal/service"
)

type WatchListController struct {
	watchListService *service.WatchListService
}

func NewWatchListController() *WatchListController {
	shoppingCartService := service.NewWatchListService(context.Background(), db.Conn)

	return &WatchListController{watchListService: &shoppingCartService}
}

func (s *WatchListController) CreateWatchList(c *gin.Context) {
	loggedUser := middleware.GetLoggedUser(c)
	_, err := s.watchListService.CreateWatchList(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	response.Entity(c, http.StatusCreated, nil)
}

func (s *WatchListController) GetWatchList(c *gin.Context) {
	loggedUser := middleware.GetLoggedUser(c)
	cart, err := s.watchListService.GetWatchList(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	response.Entity(c, http.StatusOK, cart)
}

func (s *WatchListController) ExcludeWatchList(c *gin.Context) {
	loggedUser := middleware.GetLoggedUser(c)
	err := s.watchListService.DeleteWatchList(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	response.Entity(c, http.StatusOK, nil)
}

func (s *WatchListController) AddCompanyInWatchList(c *gin.Context) {
	loggedUser := middleware.GetLoggedUser(c)

	cart, err := s.watchListService.AddCompanyInWatchListById(loggedUser, c.Param("id"))
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	response.Entity(c, http.StatusOK, cart)
}

func (s *WatchListController) AddCompanyByCodeInWatchList(c *gin.Context) {
	loggedUser := middleware.GetLoggedUser(c)

	WatchList, err := s.watchListService.AddCompanyInWatchListByCode(loggedUser, c.Param("companyCode"))
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	response.Entity(c, http.StatusOK, WatchList)
}
