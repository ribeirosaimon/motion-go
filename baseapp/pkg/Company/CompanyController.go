package Company

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	motionHttp "github.com/ribeirosaimon/motion-go/internal/httpresponse"
)

type controller struct {
	productService *Service
}

func NewCompanyController(pService *Service) controller {
	return controller{productService: pService}
}

func (c controller) saveCompany(ctx *gin.Context) {
	var companyDto companyDto

	if err := ctx.BindJSON(&companyDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	company, err := c.productService.saveCompany(companyDto)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	motionHttp.Created(ctx, company)
}

func (c controller) updateProduct(ctx *gin.Context) {

	var productDto companyDto

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}

	if err := ctx.BindJSON(&productDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	product, err := c.productService.updateCompany(productDto, id)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	motionHttp.Ok(ctx, product)
}

func (c controller) deleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}

	if !c.productService.deleteCompany(id) {
		exceptions.MotionError("cannot be deleted").Throw(ctx)
		return
	}
	motionHttp.Ok(ctx, nil)
}

func (c controller) getProduct(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	product, err := c.productService.GetCompany(id)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	motionHttp.Ok(ctx, product)
}

type companyDto struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
