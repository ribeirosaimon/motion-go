package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUser(t *testing.T) {
	// Cria um router do Gin vazio
	r := gin.New()

	NewHeathController()

	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Código HTTP de retorno esperado: %d; Código HTTP de retorno recebido: %d", http.StatusOK, w.Code)
	}

	var healthApiResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &healthApiResponse)
	if err != nil {
		t.Errorf("Erro ao decodificar o JSON de retorno: %s", err.Error())
	}
	if healthApiResponse["health"] != "true" {
		t.Errorf("Dados de usuário incorretos no JSON de retorno: %s", w.Body.String())
	}
}
