package shoppingcart

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/security"
)

type controller struct {
	shoppingCartService service
}

func NewShoppingCartController(shoppingCartService service) controller {
	return controller{shoppingCartService: shoppingCartService}
}

func (s controller) createShoppingCart(c *gin.Context) {
	loggedUser := security.GetLoggedUser(c)
	s.shoppingCartService.createShoppingCart(loggedUser)
}
