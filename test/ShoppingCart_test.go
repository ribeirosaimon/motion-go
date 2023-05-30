package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/router"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestSaveShoppingCartController(t *testing.T) {
	var e = CreateEngine(router.NewShoppingCartRouter)

	w, u := PerformRequest(e, http.MethodPost, "/shopping-cart/create", "USER", nil)

	cartRepository := repository.NewShoppingCartRepository(context.Background(), db.Conn.GetMongoTemplate())
	shopingCart, _ := cartRepository.FindByField("owner.name", u.Name)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, u.Name, shopingCart.Owner.Name)
}
