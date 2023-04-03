package test

import (
	"fmt"
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

	if resp.Code != http.StatusCreated {
		t.Errorf(util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusCreated, resp.Code)))
	}

}

func TestHaveNotCreateShoppingCartAndReturnError(t *testing.T) {
	// creating a new shopping cart
	util.AddController(testEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, MyToken, domain.USER)

	if err != nil {
		t.Errorf(util.ErrorTest("Need to return a error"))
	}

	if resp.Code != http.StatusConflict {
		t.Errorf(util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusCreated, resp.Code)))
	}
}
