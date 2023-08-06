package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/dto"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/stretchr/testify/assert"
)

func TestNewOpenHealthController(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	NewHealthController().OpenHealth(c)
	var response dto.HealthApiResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, true, response.Ready)
}

func TestNewCloseHealthControllerReturnForbidden(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	NewHealthController().CloseHealth(c)
	var response dto.HealthApiResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, false, response.Ready)
}

func TestNewCloseHealthControllerReturnOk(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("loggedUser", middleware.LoggedUser{})
	NewHealthController().CloseHealth(c)
	var response dto.HealthApiResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, false, response.Ready)
}
