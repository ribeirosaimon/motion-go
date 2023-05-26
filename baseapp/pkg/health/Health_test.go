package health

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

func TestNewOpenHealthController(t *testing.T) {
	e := test.CreateEngine(NewHealthRouter)
	w, _ := test.PerformRequest(e, http.MethodGet, "/health/open", "", nil)
	var res healthApiResponse

	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, res.Ready, true)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCloseHealthControllerError(t *testing.T) {
	e := test.CreateEngine(NewHealthRouter)
	w, _ := test.PerformRequest(e, http.MethodGet, "/health/close", "", nil)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestCloseHealthControllerSuccess(t *testing.T) {

	e := test.CreateEngine(NewHealthRouter)

	w, _ := test.PerformRequest(e, http.MethodGet, "/health/close", "USER", nil)
	var res healthApiResponse

	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, res.Ready, true)
	assert.Equal(t, http.StatusOK, w.Code)
}
