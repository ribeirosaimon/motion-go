package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/src/pkg/dto"
	"github.com/ribeirosaimon/motion-go/src/pkg/service"
	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

func TestPortfolioController_CreatePortfolio(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	NewPortfolioController().CreatePortfolio(c)
	assert.Equal(t, http.StatusOK, w.Code)
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

	var price = dto.BuyPriceDTO{
		Price:    11.5,
		Quantity: 100,
	}

	jsonBytes, err := json.Marshal(price)
	reader := bytes.NewReader(jsonBytes)

	c.Request = &http.Request{Body: ioutil.NopCloser(reader)}

	portfolioService := service.NewPortfolioService(c, db.Conn)
	portfolioService.CreatePortfolio(loggedUser)
	portfolioService.AddCompanyInPortfolioByCode(loggedUser, "TEST2", price)

	NewPortfolioController().AddCompanyByCodeInPortfolio(c)
	var response nosqlDomain.Portfolio
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	portfolioRepository := repository.NewPortfolioRepository(c, db.Conn.GetMongoTemplate())
	portfolio, _ := portfolioRepository.FindById(response.Id.Hex())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, portfolio.Id, response.Id)
	assert.Equal(t, portfolio.Companies[0].StockId, response.Companies[0].StockId)
	assert.Equal(t, portfolio.Status, response.Status)
	assert.Equal(t, portfolio.Price, float64(2300))
}

func TestPortfolioController_AddCompanyInPortfolio(t *testing.T) {
	w, c, stock1, _ := configTest()
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	var param = &gin.Params{
		{
			Key:   "id",
			Value: stock1.Id.Hex(),
		},
	}
	c.Params = *param

	var price = dto.BuyPriceDTO{
		Price:    11.5,
		Quantity: 100,
	}

	jsonBytes, err := json.Marshal(price)
	reader := bytes.NewReader(jsonBytes)

	c.Request = &http.Request{Body: ioutil.NopCloser(reader)}

	portfolioService := service.NewPortfolioService(c, db.Conn)
	portfolioService.CreatePortfolio(loggedUser)

	NewPortfolioController().AddCompanyInPortfolio(c)
	var response nosqlDomain.Portfolio
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	portfolioRepository := repository.NewPortfolioRepository(c, db.Conn.GetMongoTemplate())
	portfolio, _ := portfolioRepository.FindById(response.Id.Hex())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, portfolio.Id, response.Id)
	assert.Equal(t, portfolio.Companies[0].StockId, response.Companies[0].StockId)
	assert.Equal(t, portfolio.Status, response.Status)
	assert.Equal(t, portfolio.Price, float64(1150))
}

func TestPortfolioController_AddCompanyInPortfolioWithError(t *testing.T) {
	w, c, stock1, _ := configTest()
	test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	var param = &gin.Params{
		{
			Key:   "id",
			Value: stock1.Id.Hex(),
		},
	}
	c.Params = *param

	NewPortfolioController().AddCompanyInPortfolio(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestPortfolioController_ExcludePortfolio(t *testing.T) {
	w, c, _, _ := configTest()
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	portfolioService := service.NewPortfolioService(c, db.Conn)
	portfolioService.CreatePortfolio(loggedUser)

	NewPortfolioController().ExcludePortfolio(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPortfolioController_ExcludePortfolioDoesNotExist(t *testing.T) {
	w, c, _, _ := configTest()
	test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())

	NewPortfolioController().ExcludePortfolio(c)
	assert.Equal(t, http.StatusConflict, w.Code)
}
