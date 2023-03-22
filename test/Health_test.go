package test

import (
	"encoding/json"
	"fmt"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/security"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var healthEnginer = gin.New()

func BenchmarkController(b *testing.B) {
	start := time.Now()
	util.AddController(healthEnginer, "/api/v1/health", health.NewHealthRouter)
	resp, req, err := util.CreateEngineRequest(healthEnginer, http.MethodGet, "/api/v1/health/open",
		nil, "", domain.USER)
	if resp.Code != http.StatusOK {
		util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusOK, resp.Code))
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
	util.AddController(healthEnginer, "/api/v1/health", health.NewHealthRouter)
	resp, _, err := util.CreateEngineRequest(healthEnginer, http.MethodGet, "/api/v1/health/open",
		nil, "", domain.USER)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected Http status: %d; but is received: %d", http.StatusOK, resp.Code)
	}

	var response healthApiResponse
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Unmarshal erro: %s", err.Error())
	}
	if response.Ready != true {
		t.Errorf("Status must be bool: %s", resp.Body.String())
	}
	if &response.Time == nil {
		t.Errorf("Time must be unlike nil: %s", resp.Body.String())
	}
}

func TestCloseControllerSendError(t *testing.T) {
	t.Log("Test close controller send error")
	util.AddController(healthEnginer, "/api/v1/health", health.NewHealthRouter)
	resp, _, err := util.CreateEngineRequest(healthEnginer, http.MethodGet, "/api/v1/health/close",
		nil, "", domain.USER)
	if err != nil {
		util.ErrorTest(fmt.Sprintf("erro"))
	}
	if resp.Code != http.StatusForbidden {
		util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d",
			http.StatusForbidden, resp.Code))
	}
}

func TestCloseControllerSuccess(t *testing.T) {
	t.Log("Test close controller sucess")
	util.AddController(healthEnginer, "/api/v1/health", health.NewHealthRouter)
	session, err := util.SignUp(healthEnginer, domain.USER, domain.ADMIN, domain.USER)
	resp, _, err := util.CreateEngineRequest(healthEnginer, http.MethodGet, "/api/v1/health/close",
		nil, session, domain.USER)

	if resp.Code != http.StatusOK {
		util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusOK, resp.Code))
	}

	var response healthApiResponse
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		util.ErrorTest(fmt.Sprintf("Unmarshal erro: %s", err.Error()))
	}
	if response.Ready != true {
		util.ErrorTest(fmt.Sprintf("Status must be bool: %s", resp.Body.String()))
	}
	if &response.Time == nil {
		util.ErrorTest(fmt.Sprintf("Time must be unlike nil: %s", resp.Body.String()))
	}
}

type healthApiResponse struct {
	Ready      bool                `json:"ready"`
	Time       time.Time           `json:"time"`
	LoggedUSer security.LoggedUser `json:"loggedUser"`
}
