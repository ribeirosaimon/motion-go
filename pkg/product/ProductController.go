package product

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/exceptions"
	"github.com/shopspring/decimal"
)

type controller struct {
	productService *service
}

func NewProductController(pService *service) controller {
	return controller{productService: pService}
}

func (c controller) saveProduct(ctx *gin.Context) {
	var productDto ProductDto

	if err := ctx.BindJSON(&productDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	c.productService.saveProduct(productDto)
}

type ProductDto struct {
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
	Image string          `json:"image"`
}
