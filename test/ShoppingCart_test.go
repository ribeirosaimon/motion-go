package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var shoppingCartRouter = func(engine *gin.RouterGroup) {
	shoppingcart.NewShoppingCartRouter(engine.Group("/api/v1"), util.ConnectDatabaseTest)
}

var token string

func TestHaveToCreateShoppingCartAndReturnOk(t *testing.T) {
	defer util.RemoveDatabase()

	resp, _, err := util.CreateEngineRequest(http.MethodPost, "/api/v1/shopping-cart/create",
		nil, shoppingCartRouter, token, domain.USER)

	if err != nil {
		util.ErrorTest(err.Error())
	}

	if resp.Code != http.StatusCreated {
		util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusCreated, resp.Code))
	}

}

func TestHaveNotCreateShoppingCartAndReturnError(t *testing.T) {
	defer util.RemoveDatabase()
	// creating a new shopping cart
	TestHaveToCreateShoppingCartAndReturnOk(t)

	resp, _, err := util.CreateEngineRequest(http.MethodPost, "/api/v1/shopping-cart/create",
		nil, shoppingCartRouter, token, domain.USER)

	if err == nil {
		util.ErrorTest(err.Error())
	}

	if resp.Code == http.StatusCreated {
		util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusCreated, resp.Code))
	}
}

func init() {
	session, err := util.SignUp(domain.USER, domain.ADMIN, domain.USER)
	if err != nil {
		util.ErrorTest(err.Error())
	}
	token = session
}
