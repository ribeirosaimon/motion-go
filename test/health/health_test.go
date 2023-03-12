package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/test/util"
)

func BenchmarkController(b *testing.B) {
	resp, req, err := util.CreateEngineRequest(http.MethodGet, "/health", nil)
	if resp.Code != http.StatusOK {
		b.Errorf("Expected Http status: %d; but is received: %d", http.StatusOK, resp.Code)
	}
	if err != nil {
		b.Error("error in request")
	}
	for i := 0; i < b.N; i++ {
		// Grava a resposta HTTP
		response := httptest.NewRecorder()

		// Processa a solicitação HTTP usando o roteador do Gin
		util.GetEnginer().ServeHTTP(response, req)
	}
}

func TestController(t *testing.T) {
	resp, _, err := util.CreateEngineRequest(http.MethodGet, "/health", nil)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected Http status: %d; but is received: %d", http.StatusOK, resp.Code)
	}

	var response = health.NewHealthApiResponse()
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
