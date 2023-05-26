package shoppingcart

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

var e = test.CreateEngine(NewShoppingCartRouter)

func TestSaveShoppingCartController(t *testing.T) {

	w := test.PerformRequest(e, http.MethodPost, "/shopping-cart/create", "USER", nil)

	var res sqlDomain.Company
	json.Unmarshal([]byte(w.Body.String()), &res)

	conn := db.Conn.GetPgsqTemplate()
	defer db.Conn.ClosePostgreSQL()
	dbCompany, _ := repository.NewCompanyRepository(conn).FindById(res.Id)

	assert.NotEqual(t, 0, res.Id)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, company.Name, dbCompany.Name)
	assert.Equal(t, company.Image, dbCompany.Image)
}
