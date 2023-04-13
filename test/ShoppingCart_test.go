package test

import (
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestHaveToCreateShoppingCartAndReturnOk(t *testing.T) {
	util.AddController(testEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, TokenUser, domain.USER)

	if err != nil {
		t.Errorf("Unmarshal erro: %s", err.Error())
	}

	util.AssertEquals(t, resp.Code, http.StatusCreated)

}

func TestHaveNotCreateShoppingCartAndReturnError(t *testing.T) {
	// creating a new shopping cart
	util.AddController(testEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, TokenUser, domain.USER)

	util.AssertEquals(t, err, nil)
	util.AssertEquals(t, resp.Code, http.StatusConflict)
}

func TestHaveToDeleteShoppingCartAndReturnOk(t *testing.T) {
	//TestHaveToCreateShoppingCartAndReturnOk(t)
	util.AddController(testEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodDelete, "/api/v1/shopping-cart",
		nil, TokenUser, domain.USER)
	if err != nil {
		t.Errorf("Unmarshal erro: %s", err.Error())
	}
	util.AssertEquals(t, http.StatusOK, resp.Code)
	util.AssertEquals(t, "", resp.Body.String())
}
