package test

import (
	"context"
	"fmt"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/router"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestSavePortfolioController(t *testing.T) {
	defer db.Conn.GetMongoTemplate().Database(db.Conn.DatabaseName).Drop(context.Background())
	var e = CreateEngine(router.NewPortfolioRouter)

	w, u := PerformRequest(e, http.MethodPost, "/portfolio", "USER", "", nil)

	cartRepository := repository.NewPortfolioRepository(context.Background(), db.Conn.GetMongoTemplate())
	shopingCart, _ := cartRepository.FindByField("owner.name", u.Name)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, u.Name, shopingCart.Owner.Name)
}

func TestAddCompanyInPortfolioController(t *testing.T) {
	defer db.Conn.GetMongoTemplate().Database(db.Conn.DatabaseName).Drop(context.Background())
	var e = CreateEngine(router.NewPortfolioRouter)

	defer db.Conn.ClosePostgreSQL()
	stock := nosqlDomain.SummaryStock{CompanyCode: "test"}
	db.Conn.GetMongoTemplate().Database("Motion").Collection("summaryStock").
		InsertOne(context.Background(), stock)

	y, dto := PerformRequest(e, http.MethodPost, "/portfolio", "USER", "", nil)
	assert.Equal(t, http.StatusCreated, y.Code)
	w, _ := PerformRequest(e, http.MethodPost, fmt.Sprintf("/portfolio/company/%s", stock.CompanyCode), string(dto.loginTestDto.LoggedRole), dto.loginTestDto.Token, nil)

	cartRepository := repository.NewPortfolioRepository(context.Background(), db.Conn.GetMongoTemplate())
	shoopingCart, _ := cartRepository.FindByField("owner.name", dto.Name)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, dto.Name, shoopingCart.Owner.Name)
	assert.True(t, containsCompany(stock.Id, shoopingCart.Companies))
}

func TestAddCompanyInPortfolioControllerWithError(t *testing.T) {
	defer db.Conn.GetMongoTemplate().Database(db.Conn.DatabaseName).Drop(context.Background())
	var e = CreateEngine(router.NewPortfolioRouter)

	y, dto := PerformRequest(e, http.MethodPost, "/portfolio", "USER", "", nil)
	assert.Equal(t, http.StatusCreated, y.Code)

	w, _ := PerformRequest(e, http.MethodPost, fmt.Sprintf("/portfolio/company/%d", 9999),
		string(dto.loginTestDto.LoggedRole), dto.loginTestDto.Token, nil)
	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestNotAddCompanyInPortfolio(t *testing.T) {
	defer db.Conn.GetMongoTemplate().Database(db.Conn.DatabaseName).Drop(context.Background())
	var e = CreateEngine(router.NewPortfolioRouter)

	defer db.Conn.ClosePostgreSQL()
	stock := nosqlDomain.SummaryStock{CompanyCode: "test"}
	db.Conn.GetMongoTemplate().Database("Motion").Collection("summaryStock").
		InsertOne(context.Background(), stock)

	y, dto := PerformRequest(e, http.MethodPost, "/portfolio", "USER", "", nil)
	assert.Equal(t, http.StatusCreated, y.Code)
	PerformRequest(e, http.MethodPost, fmt.Sprintf("/portfolio/company/!!"), string(dto.loginTestDto.LoggedRole), dto.loginTestDto.Token, nil)
	w, _ := PerformRequest(e, http.MethodPost, fmt.Sprintf("/portfolio/company/!!"), string(dto.loginTestDto.LoggedRole), dto.loginTestDto.Token, nil)
	assert.Equal(t, http.StatusConflict, w.Code)
}

func containsCompany(companyId primitive.ObjectID, companies []primitive.ObjectID) bool {
	for _, v := range companies {
		if v == companyId {
			return true
		}
	}
	return false
}
