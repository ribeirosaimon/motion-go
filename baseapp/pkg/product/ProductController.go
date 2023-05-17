package product

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	motionHttp "github.com/ribeirosaimon/motion-go/internal/httpresponse"
	"github.com/shopspring/decimal"
)

type controller struct {
	productService *Service
}

func NewProductController(pService *Service) controller {
	return controller{productService: pService}
}

func (c controller) saveProduct(ctx *gin.Context) {
	var productDto ProductDto

	if err := ctx.BindJSON(&productDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	product, err := c.productService.saveProduct(productDto)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	motionHttp.Created(ctx, product)
}

func (c controller) updateProduct(ctx *gin.Context) {

	var productDto ProductDto

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}

	if err := ctx.BindJSON(&productDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	product, err := c.productService.updateProduct(productDto, id)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	motionHttp.Ok(ctx, product)
}

func (c controller) deleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("productId"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}

	if !c.productService.deleteProduct(id) {
		exceptions.MotionError("cannot be deleted").Throw(ctx)
		return
	}
	motionHttp.Ok(ctx, nil)
}

func (c controller) getProduct(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("productId"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	product, err := c.productService.GetProduct(id)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	motionHttp.Ok(ctx, product)
}

type ProductDto struct {
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
	Image string          `json:"image"`
}
