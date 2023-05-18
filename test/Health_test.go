package test

import (
	"encoding/json"
	"github.com/ribeirosaimon/motion-go/test/util"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/health"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

var (
	service    = health.NewHealthService()
	controller = health.NewHealthController(&service)
)

func TestMyOpenController(t *testing.T) {

	recorder := httptest.NewRecorder()
	context := util.GetTestGinContext(recorder)

	controller.OpenHealth(context)

	var res healthApiResponse
	json.Unmarshal([]byte(recorder.Body.String()), &res)

	assert.Equal(t, res.Ready, true)
	assert.Equal(t, recorder.Code, 200)
}

func TestCloseControllerSendError(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := util.GetTestGinContext(recorder)

	controller.CloseHealth(context)
	assert.Equal(t, recorder.Code, http.StatusForbidden)
}

func TestCloseControllerSuccess(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/health/close", nil)
	resp := httptest.NewRecorder()
	AddUserTokenInReq(req)
	TestEnginer.MotionEngine.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response healthApiResponse
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, err, nil)
	assert.Equal(t, response.Ready, true)
	assert.Equal(t, response.Time.Day(), time.Now().Day())
}

type healthApiResponse struct {
	Ready      bool                  `json:"ready"`
	Time       time.Time             `json:"time"`
	LoggedUser middleware.LoggedUser `json:"loggedUser"`
}
