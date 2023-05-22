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

	testEnginer.MotionEngine.ServeHTTP(resp, req)

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
	AddRouter(productVersion)
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

	u := &url.URL{Path: fmt.Sprintf("/api/v1/product/%d", idProduct)}
	params := url.Values{}
	params.Add("id", "1")
	u.RawQuery = params.Encode()

	Context().Request, err = http.NewRequest(http.MethodPut, u.String(), bytes.NewReader(jsonData))
	AddAdminTokenInReq(Context().Request)

	testEnginer.MotionEngine.HandleContext(Context())

	var productResponse = sqlDomain.Product{}
	json.Unmarshal(HttpResponse().Body.Bytes(), &productResponse)

	assert.Equal(t, HttpResponse().Code, http.StatusOK)
	assert.Equal(t, productResponse.Image, productDto.Image)
	assert.Equal(t, productResponse.Id, idProduct)
}

func TestHaveToPutProductAndReturnOkWithContext(t *testing.T) {
	AddRouter(productVersion)
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

	u := &url.URL{Path: fmt.Sprintf("/api/v1/product/%d", idProduct)}

	Context().Request, err = http.NewRequest(http.MethodPut, u.String(), bytes.NewReader(jsonData))
	AddAdminTokenInReq(Context().Request)

	Context().Params = []gin.Param{
		{Key: "id", Value: fmt.Sprintf("%d", idProduct)},
	}

	// sprintf := fmt.Sprintf(Context().Param("id"))
	Context().Request.URL.Path = u.String()
	// fmt.Println(sprintf, Context().Request.URL.Path)
	// testEnginer.MotionEngine.HandleContext(Context())
	myRec := httptest.NewRecorder()
	testEnginer.MotionEngine.ServeHTTP(myRec, Context().Request)

	var productResponse = sqlDomain.Product{}
	json.Unmarshal(myRec.Body.Bytes(), &productResponse)

	assert.Equal(t, HttpResponse().Code, http.StatusOK)
	assert.Equal(t, productResponse.Image, productDto.Image)
	assert.Equal(t, productResponse.Id, idProduct)
	// TODO  https://github.com/gin-gonic/gin/blob/master/routes_test.go
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
