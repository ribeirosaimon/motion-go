package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ribeirosaimon/motion-go/pkg/health"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"github.com/ribeirosaimon/motion-go/test/util"
)

func BenchmarkController(b *testing.B) {
	start := time.Now()
	util.AddController(testEnginer, "/api/v1/health", health.NewHealthRouter)
	resp, req, err := util.CreateEngineRequest(testEnginer, http.MethodGet, "/api/v1/health/open",
		nil, "", domain.USER)
	if resp.Code != http.StatusOK {
		// util.ErrorTest(b, http.StatusOK, resp.Code)
	}
	if err != nil {
		b.Error("error in request")
	}
	for i := 0; i < b.N; i++ {
		// Grava a resposta HTTP
		response := httptest.NewRecorder()

		// Processa a solicitação HTTP usando o roteador do Gin
		gin.New().ServeHTTP(response, req)
	}
	duration := time.Since(start)
	util.SuccessTest(fmt.Sprintf("Its all ok! Time: %f", float64(duration.Microseconds())/1000))
}

func TestOpenController(t *testing.T) {
	t.Log("Test open controller")
	util.AddController(testEnginer, "/api/v1/health", health.NewHealthRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodGet, "/api/v1/health/open",
		nil, "", domain.USER)

	util.ErrorTest(t, resp.Code, http.StatusOK)

	var response healthApiResponse
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	util.ErrorTest(t, nil, err)
	util.ErrorTest(t, response.Ready, true)
	util.ErrorTest(t, response.Time, nil)
}

func TestCloseControllerSendError(t *testing.T) {
	t.Log("Test close controller send error")
	util.AddController(testEnginer, "/api/v1/health", health.NewHealthRouter)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodGet, "/api/v1/health/close",
		nil, "", domain.USER)
	util.ErrorTest(t, nil, err)
	util.ErrorTest(t, http.StatusForbidden, resp.Code)
}

func TestCloseControllerSuccess(t *testing.T) {
	t.Log("Test close controller sucess")
	util.AddController(testEnginer, "/api/v1/health", health.NewHealthRouter)
	session, err := util.SignUp(testEnginer, domain.USER, domain.ADMIN, domain.USER)
	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodGet, "/api/v1/health/close",
		nil, session, domain.USER)

	util.ErrorTest(t, http.StatusOK, resp.Code)

	var response healthApiResponse
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	util.ErrorTest(t, err, nil)
	util.ErrorTest(t, response.Ready, true)
	util.ErrorTest(t, &response.Time, nil)
}

type healthApiResponse struct {
	Ready      bool                `json:"ready"`
	Time       time.Time           `json:"time"`
	LoggedUSer security.LoggedUser `json:"loggedUser"`
}
