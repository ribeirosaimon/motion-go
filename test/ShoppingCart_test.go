package test

import (
	"context"
	"fmt"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/router"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestSaveShoppingCartController(t *testing.T) {
	defer db.Conn.GetMongoTemplate().Database(db.Conn.DatabaseName).Drop(context.Background())
	var e = CreateEngine(router.NewShoppingCartRouter)

	w, u := PerformRequest(e, http.MethodPost, "/shopping-cart", "USER", "", nil)

	cartRepository := repository.NewShoppingCartRepository(context.Background(), db.Conn.GetMongoTemplate())
	shopingCart, _ := cartRepository.FindByField("owner.name", u.Name)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, u.Name, shopingCart.Owner.Name)
}

func TestAddCompanyInShoppingCartController(t *testing.T) {
	defer db.Conn.GetMongoTemplate().Database(db.Conn.DatabaseName).Drop(context.Background())
	var e = CreateEngine(router.NewShoppingCartRouter)

	defer db.Conn.ClosePostgreSQL()
	dbCompany, _ := repository.NewCompanyRepository(db.Conn.GetPgsqTemplate()).Save(createCompany())

	y, dto := PerformRequest(e, http.MethodPost, "/shopping-cart", "USER", "", nil)
	assert.Equal(t, http.StatusCreated, y.Code)
	w, u := PerformRequest(e, http.MethodPost, fmt.Sprintf("/shopping-cart/company/%d", dbCompany.Id), string(dto.loginTestDto.LoggedRole), dto.loginTestDto.Token, nil)

	cartRepository := repository.NewShoppingCartRepository(context.Background(), db.Conn.GetMongoTemplate())
	shoopingCart, _ := cartRepository.FindByField("owner.name", dto.Name)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, u.Name, shoopingCart.Owner.Name)
	assert.True(t, containsCompany(dbCompany.Name, shoopingCart.Companies))
}

func TestAddCompanyInShoppingCartControllerWithError(t *testing.T) {
	defer db.Conn.GetMongoTemplate().Database(db.Conn.DatabaseName).Drop(context.Background())
	var e = CreateEngine(router.NewShoppingCartRouter)

	y, dto := PerformRequest(e, http.MethodPost, "/shopping-cart", "USER", "", nil)
	assert.Equal(t, http.StatusCreated, y.Code)

	w, _ := PerformRequest(e, http.MethodPost, fmt.Sprintf("/shopping-cart/company/%d", 9999),
		string(dto.loginTestDto.LoggedRole), dto.loginTestDto.Token, nil)
	assert.Equal(t, http.StatusConflict, w.Code)
}

func containsCompany(companyName string, companies []sqlDomain.Company) bool {
	for _, v := range companies {
		if v.Name == companyName {
			return true
		}
	}
	return false
}
