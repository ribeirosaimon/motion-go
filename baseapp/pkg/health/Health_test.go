package health

import (
	"encoding/json"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

func TestNewOpenHealthController(t *testing.T) {
	e := test.CreateEngine(NewHealthRouter)
	w := test.PerformRequest(e, http.MethodGet, "/health/open", nil, nil)
	var res healthApiResponse

	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, res.Ready, true)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCloseHealthControllerError(t *testing.T) {
	e := test.CreateEngine(NewHealthRouter)
	w := test.PerformRequest(e, http.MethodGet, "/health/close", nil, nil)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestCloseHealthControllerSuccess(t *testing.T) {

	e := test.CreateEngine(NewHealthRouter)

	userRole := sqlDomain.USER
	w := test.PerformRequest(e, http.MethodGet, "/health/close", &userRole, nil)
	var res healthApiResponse

	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, res.Ready, true)
	assert.Equal(t, http.StatusOK, w.Code)
}
