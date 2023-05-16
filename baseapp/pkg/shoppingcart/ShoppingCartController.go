package shoppingcart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

type controller struct {
	shoppingCartService *service
}

func NewShoppingCartController(shoppingCartService *service) controller {
	return controller{shoppingCartService: shoppingCartService}
}

func (s controller) createShoppingCart(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	_, err = s.shoppingCartService.createShoppingCart(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	c.Status(http.StatusCreated)
}

func (s controller) getShoppingCart(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	cart, err := s.shoppingCartService.getShoppingCart(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Created(c, cart)
}

func (s controller) excludeShoppingCart(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	err = s.shoppingCartService.deleteShoppingCart(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Ok(c, nil)
}

func (s controller) addProductInShoppingCart(c *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(c)
	var productDTO productDTO

	if err := c.BindJSON(&productDTO); err != nil {
		exceptions.BodyError().Throw(c)
		return
	}

	cart, err := s.shoppingCartService.addProductInShoppingCart(loggedUser, productDTO)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Ok(c, cart)
}

type productDTO struct {
	Id     int64 `json:"id"`
	Amount int64 `json:"amount"`
}
