package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/product"
	"github.com/ribeirosaimon/motion-go/test/util"
	"github.com/shopspring/decimal"
)

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
	util.ErrorTest(t, resp.Code, http.StatusOK)
	util.ErrorTest(t, productResponse.Image, productDto.Image)
}
