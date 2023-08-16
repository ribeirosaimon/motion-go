package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPortfolioController_CreatePortfolio(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	NewPortfolioController().CreatePortfolio(c)
	assert.Equal(t, http.StatusCreated, w.Code)
}
