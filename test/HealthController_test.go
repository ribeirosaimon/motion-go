package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ribeirosaimon/motion-go/test/util"

	"github.com/magiconair/properties/assert"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/health"
)

var (
	service    = health.NewHealthService()
	controller = health.NewHealthController(&service)
)

func TestMyOpenController(t *testing.T) {

	recorderContext := httptest.NewRecorder()
	context := util.GetTestGinContext(recorderContext)

	controller.OpenHealth(context)

	var res healthApiResponse
	json.Unmarshal([]byte(recorderContext.Body.String()), &res)

	assert.Equal(t, res.Ready, true)
	assert.Equal(t, recorderContext.Code, 200)
}

func TestCloseControllerSuccess(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/health/close", nil)
	recorderContext := httptest.NewRecorder()
	AddUserTokenInReq(req)
	testEnginer.MotionEngine.ServeHTTP(recorderContext, req)

	assert.Equal(t, http.StatusOK, recorderContext.Code)

	var response healthApiResponse
	err := json.Unmarshal(recorderContext.Body.Bytes(), &response)
	assert.Equal(t, err, nil)
	assert.Equal(t, response.Ready, true)
	assert.Equal(t, response.Time.Day(), time.Now().Day())
}
