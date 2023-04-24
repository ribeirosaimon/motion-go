package test

import (
	"github.com/magiconair/properties/assert"
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"

	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"

	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestHaveToCreateShoppingCartAndReturnOk(t *testing.T) {
	util.AddController(testEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, TokenUser, sqlDomain.USER)

	if err != nil {
		t.Errorf("Unmarshal erro: %s", err.Error())
	}

	assert.Equal(t, resp.Code, http.StatusCreated)

}

func TestHaveNotCreateShoppingCartAndReturnError(t *testing.T) {
	// creating a new shopping cart
	util.AddController(testEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPost, "/api/v1/shopping-cart/create",
		nil, TokenUser, sqlDomain.USER)

	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, http.StatusConflict)
}

func TestHaveToDeleteShoppingCartAndReturnOk(t *testing.T) {
	// TestHaveToCreateShoppingCartAndReturnOk(t)
	util.AddController(testEnginer, "/api/v1/shopping-cart", shoppingcart.NewShoppingCartRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodDelete, "/api/v1/shopping-cart",
		nil, TokenUser, sqlDomain.USER)
	if err != nil {
		t.Errorf("Unmarshal erro: %s", err.Error())
	}
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "", resp.Body.String())
}
