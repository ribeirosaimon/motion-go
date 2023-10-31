package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/repository"
	"github.com/ribeirosaimon/motion-go/api/internal/service"
	"github.com/ribeirosaimon/motion-go/api/test"
	"github.com/ribeirosaimon/motion-go/confighub/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/confighub/domain/sqlDomain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestWatchListController_CreateWatchList(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	NewWatchListController().CreateWatchList(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestWatchListController_CreateWatchListReturnError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())

	WatchListService := service.NewWatchListService(c, db.Conn)
	WatchListService.CreateWatchList(loggedUser)

	NewWatchListController().CreateWatchList(c)

	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestWatchListController_GetWatchList(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())

	WatchListService := service.NewWatchListService(c, db.Conn)
	WatchListService.CreateWatchList(loggedUser)

	NewWatchListController().GetWatchList(c)

	WatchList, _ := WatchListService.GetWatchList(loggedUser)
	var response nosqlDomain.WatchList
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, WatchList.Id, response.Id)
	assert.Equal(t, WatchList.OwnerId, response.OwnerId)
	assert.Equal(t, WatchList.OwnerId, loggedUser.ProfileId)
}

func TestWatchListController_AddCompanyByCodeInWatchList(t *testing.T) {
	w, c, _, stock1, _ := configTest()
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	var param = &gin.Params{
		{
			Key:   "companyCode",
			Value: stock1.CompanyCode,
		},
	}
	c.Params = *param

	WatchListService := service.NewWatchListService(c, db.Conn)
	WatchListService.CreateWatchList(loggedUser)
	WatchListService.AddCompanyInWatchListByCode(loggedUser, "TEST2")

	NewWatchListController().AddCompanyByCodeInWatchList(c)
	var response nosqlDomain.WatchList
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	WatchListRepository := repository.NewWatchListRepository(c, db.Conn.GetMongoTemplate())
	WatchList, _ := WatchListRepository.FindById(response.Id.Hex())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, WatchList.Id, response.Id)
	assert.Equal(t, WatchList.Companies[0], response.Companies[0])
}

func TestWatchListController_AddCompanyInWatchList(t *testing.T) {
	w, c, _, stock1, _ := configTest()
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	var param = &gin.Params{
		{
			Key:   "id",
			Value: stock1.Id.Hex(),
		},
	}
	c.Params = *param

	WatchListService := service.NewWatchListService(c, db.Conn)
	WatchListService.CreateWatchList(loggedUser)

	NewWatchListController().AddCompanyInWatchList(c)
	var response nosqlDomain.WatchList
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	WatchListRepository := repository.NewWatchListRepository(c, db.Conn.GetMongoTemplate())
	WatchList, _ := WatchListRepository.FindById(response.Id.Hex())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, WatchList.Id, response.Id)
	assert.Equal(t, WatchList.Companies[0], response.Companies[0])
}

func TestWatchListController_AddCompanyInWatchListWithError(t *testing.T) {
	w, c, _, stock1, _ := configTest()
	test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	var param = &gin.Params{
		{
			Key:   "id",
			Value: stock1.Id.Hex(),
		},
	}
	c.Params = *param

	NewWatchListController().AddCompanyInWatchList(c)

	assert.Equal(t, http.StatusConflict, w.Code)

}

func TestWatchListController_ExcludeWatchList(t *testing.T) {
	w, c, _, _, _ := configTest()
	loggedUser := test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())
	WatchListService := service.NewWatchListService(c, db.Conn)
	WatchListService.CreateWatchList(loggedUser)

	NewWatchListController().ExcludeWatchList(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestWatchListController_ExcludeWatchListDoesNotExist(t *testing.T) {
	w, c, _, _, _ := configTest()
	test.SetUpTest(c, sqlDomain.USER)
	defer db.Conn.GetMongoTemplate().Database(db.Conn.GetMongoDatabase()).Drop(context.Background())

	NewWatchListController().ExcludeWatchList(c)
	assert.Equal(t, http.StatusConflict, w.Code)
}
