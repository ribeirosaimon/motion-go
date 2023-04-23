package shoppingcart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
	"github.com/ribeirosaimon/motion-go/pkg/security"
)

type controller struct {
	shoppingCartService *service
}

func NewShoppingCartController(shoppingCartService *service) controller {
	return controller{shoppingCartService: shoppingCartService}
}

func (s controller) createShoppingCart(c *gin.Context) {
	loggedUser := security.GetLoggedUser(c)
	_, err := s.shoppingCartService.createShoppingCart(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	c.Status(http.StatusCreated)
}

func (s controller) getShoppingCart(c *gin.Context) {
	loggedUser := security.GetLoggedUser(c)
	cart, err := s.shoppingCartService.getShoppingCart(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Created(c, cart)
}

func (s controller) excludeShoppingCart(c *gin.Context) {
	loggedUser := security.GetLoggedUser(c)
	err := s.shoppingCartService.deleteShoppingCart(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(c)
		return
	}
	httpresponse.Ok(c, nil)
}

func (s controller) addProductInShoppingCart(c *gin.Context) {
	// loggedUser := security.GetLoggedUser(c) //TODO

}
