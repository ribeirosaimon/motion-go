package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/product"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/shopspring/decimal"
)

var idProduct uint64

var productVersion = config.RoutersVersion{
	Version: "v1",
	Handlers: []func() config.MotionController{
		product.NewProductRouter,
	},
}

func TestHaveToAddProductAndReturnOk(t *testing.T) {
	AddRouter(productVersion)

	bd1, _ := decimal.NewFromString("123.456")
	productDto := product.ProductDto{
		Price: bd1,
		Name:  "Teste",
		Image: "http://teste",
	}
	jsonData, err := json.Marshal(productDto)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/product", bytes.NewReader(jsonData))
	AddAdminTokenInReq(req)
	resp := httptest.NewRecorder()

	TestEnginer.MotionEngine.ServeHTTP(resp, req)

	if err != nil {
		t.Error("error in request")
	}
	var productResponse = sqlDomain.Product{}
	err = json.Unmarshal(resp.Body.Bytes(), &productResponse)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, resp.Code, http.StatusCreated)
	assert.Equal(t, productResponse.Image, productDto.Image)
	idProduct = productResponse.Id
}

func TestHaveToPutProductAndReturnOk(t *testing.T) {
	if idProduct == 0 {
		TestHaveToAddProductAndReturnOk(t)
	}

	bd1, _ := decimal.NewFromString("321.61")
	productDto := product.ProductDto{
		Price: bd1,
		Name:  "TesteUpdate",
		Image: "http://update",
	}
	jsonData, err := json.Marshal(productDto)

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/product/%d", idProduct), bytes.NewReader(jsonData))
	AddAdminTokenInReq(req)

	resp := httptest.NewRecorder()

	params := []gin.Param{
		{
			Key:   "id",
			Value: fmt.Sprintf("%d", idProduct),
		},
	}
	//ginCtx := util.GetTestGinContext(resp)
	fmt.Sprintf("%s", params)
	//ginCtx.Set("id", idProduct)
	//ginCtx.Params = params
	u := url.Values{}
	u.Add("foo", "bar")

	//ginCtx.Request.URL.RawQuery = u.Encode()
	// service := product.NewProductService(db.Conn)
	// controller := product.NewProductController(&service)
	// controller
	//context := req.WithContext(ginCtx)
	//TestEnginer.MotionEngine.ServeHTTP(resp, context)

	var productResponse = sqlDomain.Product{}
	err = json.Unmarshal(resp.Body.Bytes(), &productResponse)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, resp.Code, http.StatusOK)
	assert.Equal(t, productResponse.Image, productDto.Image)
	assert.Equal(t, productResponse.Id, idProduct)
}

// func TestHaveToGetProductAndReturnOk(t *testing.T) {
// 	TestHaveToAddProductAndReturnOk(t)
// 	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodGet,
// 		fmt.Sprintf("/api/v1/product/%d", idProduct),
// 		nil, TokenAdmin, sqlDomain2.ADMIN)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(t, resp.Code, http.StatusOK)
// }
//
// func TestHaveToDeleteProductAndReturnOk(t *testing.T) {
// 	TestHaveToAddProductAndReturnOk(t)
// 	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodDelete,
// 		fmt.Sprintf("/api/v1/product/%d", idProduct),
// 		nil, TokenAdmin, sqlDomain2.ADMIN)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(t, resp.Code, http.StatusOK)
// }
//
// func TestCanotSaveSameProduct(t *testing.T) {
// 	TestHaveToAddProductAndReturnOk(t)
// 	bd1, _ := decimal.NewFromString("3.46")
// 	productDto := product2.ProductDto{
// 		Price: bd1,
// 		Name:  "Teste",
// 		Image: "http://teste",
// 	}
// 	jsonData, err := json.Marshal(productDto)
//
// 	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPost, "/api/v1/product",
// 		bytes.NewReader(jsonData), TokenAdmin, sqlDomain2.ADMIN)
//
// 	if err != nil {
// 		t.Error("error in request")
// 	}
// 	assert.Equal(t, resp.Code, http.StatusConflict)
// }
