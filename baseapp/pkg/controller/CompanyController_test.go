package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCompanyController_GetCompany_ReturnOk(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	NewCompanyController()
	var response dto.HealthApiResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, true, response.Ready)
}
