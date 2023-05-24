package Company

import (
	"bytes"
	"encoding/json"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestProductController(t *testing.T) {
	e := test.CreateEngine(NewCompanyRouter)

	var company sqlDomain.Company

	company.Name = "petro"
	company.Image = "testeImage"
	jsonData, _ := json.Marshal(company)
	w := test.PerformRequest(e, http.MethodPost, "/company", "ADMIN", bytes.NewReader(jsonData))

	var res sqlDomain.Company
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.NotEqual(t, 0, res.Id)
	assert.Equal(t, http.StatusOK, w.Code)
}
