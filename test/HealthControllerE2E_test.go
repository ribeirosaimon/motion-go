package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/health"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

var healthVersion = config.RoutersVersion{
	Version: "v1",
	Handlers: []func() config.MotionController{
		health.NewHealthRouter,
	},
}

func TestOpen(t *testing.T) {

	AddRouter(healthVersion)
	u := &url.URL{Path: "/api/v1/health/open"}

	Context().Request, _ = http.NewRequest(http.MethodGet, u.String(), nil)

	testEnginer.MotionEngine.HandleContext(Context())

	var res healthApiResponse

	json.Unmarshal([]byte(HttpResponse().Body.String()), &res)

	assert.Equal(t, res.Ready, true)
	assert.Equal(t, HttpResponse().Code, 200)
}

func TestClose(t *testing.T) {
	AddRouter(healthVersion)
	u := &url.URL{Path: "/api/v1/health/close"}

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	resp := httptest.NewRecorder()
	testEnginer.MotionEngine.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusForbidden)
}

func TestCloseSuccess(t *testing.T) {
	AddRouter(healthVersion)
	u := &url.URL{Path: "/api/v1/health/close"}

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	AddAdminTokenInReq(req)

	resp := httptest.NewRecorder()
	testEnginer.MotionEngine.ServeHTTP(resp, req)

	var response healthApiResponse
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, err, nil)
	assert.Equal(t, response.Ready, true)
	assert.Equal(t, response.Time.Day(), time.Now().Day())
}

type healthApiResponse struct {
	Ready      bool                  `json:"ready"`
	Time       time.Time             `json:"time"`
	LoggedUser middleware.LoggedUser `json:"loggedUser"`
}
