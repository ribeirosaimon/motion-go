package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var token string

enginer = gin.
func TestHaveToCreateShoppingCartAndReturnOk(t *testing.T) {

	resp, _, err := util.CreateEngineRequest(enginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, token, domain.USER)

	if err != nil {
		util.ErrorTest(err.Error())
	}

	if resp.Code != http.StatusCreated {
		util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusCreated, resp.Code))
	}

}

func TestHaveNotCreateShoppingCartAndReturnError(t *testing.T) {
	// creating a new shopping cart
	TestHaveToCreateShoppingCartAndReturnOk(t)

	resp, _, err := util.CreateEngineRequest(enginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, token, domain.USER)

	if err == nil {
		util.ErrorTest("Need to return a error")
	}

	if resp.Code != http.StatusConflict {
		util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusCreated, resp.Code))
	}
}

func init() {
	shoppingcart.NewShoppingCartRouter(enginer.Group("/api/v1"), util.ConnectDatabaseTest)
	session, err := util.SignUp(enginer, domain.USER, domain.ADMIN, domain.USER)
	if err != nil {
		util.ErrorTest(err.Error())
	}
	token = session
}
