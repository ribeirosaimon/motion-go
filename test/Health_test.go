package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/health"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"gorm.io/gorm/utils"
)

var healthVersion = config.RoutersVersion{
	Version: "v1",
	Handlers: []func() config.MotionController{
		health.NewHealthRouter,
	},
}

func BenchmarkController(b *testing.B) {

	AddRouter(healthVersion)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/health/open", nil)

	resp := httptest.NewRecorder()

	TestEnginer.MotionEngine.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		utils.AssertEqual(http.StatusOK, resp.Code)
	}
	if err != nil {
		b.Error("error in request")
	}
	for i := 0; i < b.N; i++ {
		// Grava a resposta HTTP
		response := httptest.NewRecorder()

		// Processa a solicitação HTTP usando o roteador do Gin

		TestEnginer.MotionEngine.ServeHTTP(response, req)

	}

}

func TestOpenController(t *testing.T) {
	AddRouter(healthVersion)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/health/open", nil)

	resp := httptest.NewRecorder()

	TestEnginer.MotionEngine.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)

	var response healthApiResponse
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, nil, err)
	assert.Equal(t, response.Ready, true)
	assert.Equal(t, response.Time.Day(), time.Now().Day())
}

func TestCloseControllerSendError(t *testing.T) {
	AddRouter(healthVersion)
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/health/close", nil)
	resp := httptest.NewRecorder()

	TestEnginer.MotionEngine.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusForbidden)
}

func TestCloseControllerSuccess(t *testing.T) {
	AddRouter(healthVersion)
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
