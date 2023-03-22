package test

import (
	"fmt"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var shoppingCartEnginer = gin.New()

var myToken string

func getToken() string {

	token, err := util.SignUp(shoppingCartEnginer, domain.USER, domain.ADMIN, domain.USER)
	if err != nil {
		util.ErrorTest(err.Error())
	}

	return token
}

func TestHaveToCreateShoppingCartAndReturnOk(t *testing.T) {
	util.AddController(shoppingCartEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(shoppingCartEnginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, myToken, domain.USER)

	if err != nil {
		t.Errorf("Unmarshal erro: %s", err.Error())
	}

	if resp.Code != http.StatusCreated {
		t.Errorf(util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusCreated, resp.Code)))
	}

}

func TestHaveNotCreateShoppingCartAndReturnError(t *testing.T) {
	// creating a new shopping cart
	util.AddController(shoppingCartEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(shoppingCartEnginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, myToken, domain.USER)

	if err != nil {
		t.Errorf(util.ErrorTest("Need to return a error"))
	}

	if resp.Code != http.StatusConflict {
		t.Errorf(util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusCreated, resp.Code)))
	}
}

func init() {
	myToken = getToken()
}
