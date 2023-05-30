package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/dto"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/router"
	"github.com/stretchr/testify/assert"
)

func TestNewOpenHealthController(t *testing.T) {

	var e = CreateEngine(router.NewHealthRouter)

	w, _ := PerformRequest(e, http.MethodGet, "/health/open", "", nil)
	var res dto.HealthApiResponseDTO

	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, res.Ready, true)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCloseHealthControllerError(t *testing.T) {
	var e = CreateEngine(router.NewHealthRouter)

	w, _ := PerformRequest(e, http.MethodGet, "/health/close", "", nil)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestCloseHealthControllerSuccess(t *testing.T) {
	var e = CreateEngine(router.NewHealthRouter)

	w, _ := PerformRequest(e, http.MethodGet, "/health/close", "USER", nil)
	var res dto.HealthApiResponseDTO

	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, res.Ready, true)
	assert.Equal(t, http.StatusOK, w.Code)
}
