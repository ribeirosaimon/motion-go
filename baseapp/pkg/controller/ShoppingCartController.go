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

type shoppingCartController struct {
	shoppingCartService *service.ShoppingCartService
}

func NewShoppingCartController() shoppingCartController {
	shoppingCartService := service.NewShoppingCartService(context.Background(), db.Conn)

	return shoppingCartController{shoppingCartService: &shoppingCartService}
}

func (s shoppingCartController) CreateShoppingCart(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	_, err = s.shoppingCartService.CreateShoppingCart(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	c.Status(http.StatusCreated)
}

func (s shoppingCartController) GetShoppingCart(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	cart, err := s.shoppingCartService.GetShoppingCart(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Created(c, cart)
}

func (s shoppingCartController) ExcludeShoppingCart(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	err = s.shoppingCartService.DeleteShoppingCart(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Ok(c, nil)
}

func (s shoppingCartController) AddCompanyInShoppingCart(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(c)
		return
	}

	cart, err := s.shoppingCartService.AddProductInShoppingCart(loggedUser, id)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Ok(c, cart)
}
