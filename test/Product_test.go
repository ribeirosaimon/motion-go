package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/product"
	"github.com/ribeirosaimon/motion-go/test/util"
	"github.com/shopspring/decimal"
)

var idProduct uint64

func TestHaveToAddProductAndReturnOk(t *testing.T) {
	util.AddController(testEnginer, "/api/v1/product", product.NewProductRouter)
	bd1, _ := decimal.NewFromString("123.456")
	productDto := product.ProductDto{
		Price: bd1,
		Name:  "Teste",
		Image: "http://teste",
	}
	jsonData, err := json.Marshal(productDto)

	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPost, "/api/v1/product",
		bytes.NewReader(jsonData), TokenAdmin, domain.ADMIN)

	if err != nil {
		t.Error("error in request")
	}
	var productResponse = domain.Product{}
	err = json.Unmarshal(resp.Body.Bytes(), &productResponse)
	if err != nil {
		panic(err)
	}
	util.AssertEquals(t, resp.Code, http.StatusCreated)
	util.AssertEquals(t, productResponse.Image, productDto.Image)
	util.AssertNotEquals(t, &productResponse.Id, nil)
	idProduct = productResponse.Id
}

func TestHaveToPutProductAndReturnOk(t *testing.T) {
	TestHaveToAddProductAndReturnOk(t)

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

	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPut,
		fmt.Sprintf("/api/v1/product/%d", idProduct),
		bytes.NewReader(jsonData), TokenAdmin, domain.ADMIN)

	if err != nil {
		t.Error("error in request")
	}

	var productResponse = domain.Product{}
	err = json.Unmarshal(resp.Body.Bytes(), &productResponse)
	if err != nil {
		panic(err)
	}
	util.AssertEquals(t, resp.Code, http.StatusOK)
	util.AssertEquals(t, productResponse.Image, productDto.Image)
	util.AssertEquals(t, productResponse.Id, idProduct)

}


func TestHaveToDeleteProductAndReturnOk(t *testing.T) {
	TestHaveToAddProductAndReturnOk(t)
	.
}
