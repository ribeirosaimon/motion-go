package util

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/health"
)

var enginer *gin.Engine

func CreateEngineRequest(method, path string, body io.Reader) (
	*httptest.ResponseRecorder, *http.Request, error) {
	enginer = gin.New()

	health.NewHeathController(enginer)

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, nil, err
	}
	w := httptest.NewRecorder()
	enginer.ServeHTTP(w, req)

	return w, req, nil
}

func GetEnginer() *gin.Engine {
	return enginer
}
