package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/service"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

func TestPortfolioController_CreatePortfolio(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	NewPortfolioController().CreatePortfolio(c)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPortfolioController_CreatePortfolioReturnError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())

	portfolioService := service.NewPortfolioService(c, db.Conn)
	portfolioService.CreatePortfolio(loggedUser)

	NewPortfolioController().CreatePortfolio(c)

	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestPortfolioController_GetPortfolio(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())

	portfolioService := service.NewPortfolioService(c, db.Conn)
	portfolioService.CreatePortfolio(loggedUser)

	NewPortfolioController().GetPortfolio(c)

	portfolio, _ := portfolioService.GetPortfolio(loggedUser)
	var response nosqlDomain.Portfolio
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, portfolio.Id, response.Id)
	assert.Equal(t, portfolio.OwnerId, response.OwnerId)
	assert.Equal(t, portfolio.OwnerId, loggedUser.UserId)
	assert.Equal(t, portfolio.Status, domain.ACTIVE)
}

func TestPortfolioController_AddCompanyByCodeInPortfolio(t *testing.T) {
	w, c, stock1, _ := configTest()
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	var param = &gin.Params{
		{
			Key:   "companyCode",
			Value: stock1.CompanyCode,
		},
	}
	c.Params = *param
	portfolioService := service.NewPortfolioService(c, db.Conn)
	portfolioService.CreatePortfolio(loggedUser)

	NewPortfolioController().AddCompanyByCodeInPortfolio(c)
	var response nosqlDomain.Portfolio
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	portfolioRepository := repository.NewPortfolioRepository(c, db.Conn.GetMongoTemplate())
	portfolio, _ := portfolioRepository.FindById(response.Id.Hex())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, portfolio.Id, response.Id)
	assert.Equal(t, portfolio.Companies[0].Hex(), response.Companies[0].Hex())
	assert.Equal(t, portfolio.Status, response.Status)
}

func TestPortfolioController_AddCompanyInPortfolio(t *testing.T) {
	assert.Equal(t, false, true)
}

func TestPortfolioController_ExcludePortfolio(t *testing.T) {
	assert.Equal(t, false, true)
}
