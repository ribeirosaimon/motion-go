package test

var idProduct uint64

//
// func TestHaveToAddProductAndReturnOk(t *testing.T) {
// 	util.AddController(testEnginer, "/api/v1/product", product2.NewProductRouter)
// 	bd1, _ := decimal.NewFromString("123.456")
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
// 	var productResponse = sqlDomain2.Product{}
// 	err = json.Unmarshal(resp.Body.Bytes(), &productResponse)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(t, resp.Code, http.StatusCreated)
// 	assert.Equal(t, productResponse.Image, productDto.Image)
// 	idProduct = productResponse.Id
// }
//
// func TestHaveToPutProductAndReturnOk(t *testing.T) {
// 	TestHaveToAddProductAndReturnOk(t)
//
// 	bd1, _ := decimal.NewFromString("321.61")
// 	productDto := product2.ProductDto{
// 		Price: bd1,
// 		Name:  "TesteUpdate",
// 		Image: "http://update",
// 	}
// 	jsonData, err := json.Marshal(productDto)
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	resp, _, err := util.CreateEngineRequest(testEnginer, http.MethodPut,
// 		fmt.Sprintf("/api/v1/product/%d", idProduct),
// 		bytes.NewReader(jsonData), TokenAdmin, sqlDomain2.ADMIN)
//
// 	if err != nil {
// 		t.Error("error in request")
// 	}
//
// 	var productResponse = sqlDomain2.Product{}
// 	err = json.Unmarshal(resp.Body.Bytes(), &productResponse)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(t, resp.Code, http.StatusOK)
// 	assert.Equal(t, productResponse.Image, productDto.Image)
// 	assert.Equal(t, productResponse.Id, idProduct)
// }
//
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
