package shoppingcart

import (
	"context"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

var e = test.CreateEngine(NewShoppingCartRouter)

func TestSaveShoppingCartController(t *testing.T) {

	w, u := test.PerformRequest(e, http.MethodPost, "/shopping-cart/create", "USER", nil)

	cartRepository := repository.NewShoppingCartRepository(context.Background(), db.Conn.GetMongoTemplate())
	shopingCart, _ := cartRepository.FindByField("owner.name", u.Name)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, u.Name, shopingCart.Owner.Name)
}
