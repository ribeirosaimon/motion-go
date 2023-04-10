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
		nil, MyToken, domain.USER)

	if err != nil {
		t.Errorf("Unmarshal erro: %s", err.Error())
	}

	util.ErrorTest(t, http.StatusCreated, resp.Code)

}

func TestHaveNotCreateShoppingCartAndReturnError(t *testing.T) {
	// creating a new shopping cart
	util.AddController(testEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, MyToken, domain.USER)

	util.ErrorTest(t, err, nil)
	util.ErrorTest(t, http.StatusCreated, resp.Code)
}

func TestHaveToDeleteShoppingCartAndReturnOk(t *testing.T) {
	TestHaveToCreateShoppingCartAndReturnOk(t)
	util.AddController(testEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodDelete, "/api/v1/shopping-cart",
		nil, MyToken, domain.USER)
	if err != nil {
		t.Errorf("Unmarshal erro: %s", err.Error())
	}
	util.ErrorTest(t, http.StatusOK, resp.Code)
	util.ErrorTest(t, "", resp.Body.String())
}
