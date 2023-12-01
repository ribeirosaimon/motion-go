package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/middleware"
	"github.com/ribeirosaimon/motion-go/api/internal/repository"
	"github.com/ribeirosaimon/motion-go/api/test"
	"github.com/ribeirosaimon/motion-go/confighub/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/confighub/domain/sqlDomain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCompanyController_GetAllCompany_ReturnOk(t *testing.T) {
	w, c, _, stock1, stock2 := configTest()
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())

	NewCompanyController().GetAllCompany(c)
	var response []nosqlDomain.SummaryStock
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}
	for _, value := range response {
		if value.CompanyName == stock1.CompanyName {
			assert.Equal(t, stock1.Id.Hex(), value.Id.Hex())
		} else {

			assert.Equal(t, stock2.Id, value.Id)
		}
	}
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCompanyController_GetCompany(t *testing.T) {
	w, c, _, stock1, _ := configTest()
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	var param = &gin.Params{
		{
			Key:   "companyCode",
			Value: "test1",
		},
	}
	c.Params = *param

	NewCompanyController().GetCompanyInfo(c)

	var response nosqlDomain.SummaryStock
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}
	assert.Equal(t, stock1.Id.Hex(), response.Id.Hex())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCompanyController_DeleteCompany(t *testing.T) {
	w, c, _, stock1, stock2 := configTest()
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	var param = &gin.Params{
		{
			Key:   "id",
			Value: stock2.Id.Hex(),
		},
	}
	c.Params = *param
	NewCompanyController().DeleteCompany(c)
	stockRepository := repository.NewSummaryStockRepository(c, db.Conn.GetMongoTemplate())
	foundCompany1, err := stockRepository.FindById(stock1.Id.Hex())
	foundCompany, err := stockRepository.FindById(stock2.Id.Hex())
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, foundCompany1.Status, nosqlDomain.ACTIVE)
	assert.Equal(t, foundCompany.Status, nosqlDomain.ACTIVE)
}

func saveSummaryStock() (nosqlDomain.SummaryStock, nosqlDomain.SummaryStock) {

	var stock1 = nosqlDomain.SummaryStock{CompanyCode: "TEST1"}
	var stock2 = nosqlDomain.SummaryStock{CompanyCode: "TEST2"}
	var err error

	stockRepository := repository.NewSummaryStockRepository(context.Background(),
		db.Conn.GetMongoTemplate())
	if !stockRepository.ExistByField("companyCode", stock1.CompanyCode) {
		stock1 = nosqlDomain.SummaryStock{
			Id:          primitive.NewObjectID(),
			CompanyName: "test1",
			CompanyCode: "TEST1",
			Status:      nosqlDomain.ACTIVE,
		}

		stock1, err = stockRepository.Save(stock1)
		if err != nil {
			panic(err)
		}
	} else {
		stock1, err = stockRepository.FindByField("companyCode", stock1.CompanyCode)
		if err != nil {
			panic(err)
		}
	}
	if !stockRepository.ExistByField("companyCode", stock2.CompanyCode) {
		stock2 = nosqlDomain.SummaryStock{
			Id:          primitive.NewObjectID(),
			CompanyName: "test2",
			CompanyCode: "TEST2",
			Status:      nosqlDomain.ACTIVE,
		}

		stock2, err = stockRepository.Save(stock2)
		if err != nil {
			panic(err)
		}
	} else {
		stock2, err = stockRepository.FindByField("companyCode", stock2.CompanyCode)
		if err != nil {
			panic(err)
		}
	}
	return stock1, stock2
}

func configTest() (*httptest.ResponseRecorder, *gin.Context, middleware.LoggedUser, nosqlDomain.SummaryStock, nosqlDomain.SummaryStock) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	loggedUser := test.SetUpTest(c, sqlDomain.USER)

	stock1, stock2 := saveSummaryStock()
	middleware.Cache.Add(stock1)
	middleware.Cache.Add(stock2)
	return w, c, loggedUser, stock1, stock2
}
