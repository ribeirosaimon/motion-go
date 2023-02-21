package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/health"
)

func TestGetUser(t *testing.T) {
	// Cria um router do Gin vazio
	r := gin.New()

	health.NewHeathController(r)

	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected Http status: %d; but is received: %d", http.StatusOK, w.Code)
	}

	var response = health.NewHealthApiResponse()
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Unmarshal erro: %s", err.Error())
	}
	if response.Ready != true {
		t.Errorf("Status must be bool: %s", w.Body.String())
	}
	if &response.Time == nil {
		t.Errorf("Time must be unlike nil: %s", w.Body.String())
	}
}
