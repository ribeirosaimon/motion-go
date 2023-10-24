package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/dto"
	"github.com/ribeirosaimon/motion-go/test"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthController_OpenHealth(t *testing.T) {
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

func TestHealthController_CloseHealth_Forbidden(t *testing.T) {
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

func TestHealthController_CloseHealth_ReturnOk(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	NewHealthController().CloseHealth(c)

	var response dto.HealthApiResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, true, response.Ready)
	assert.Equal(t, loggedUser.ProfileId, response.LoggedUser.ProfileId)
}
