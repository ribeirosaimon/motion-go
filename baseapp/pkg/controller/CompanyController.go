package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/dto"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/service"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
)

type companyController struct {
	companyService *service.CompanyService
}

func NewCompanyController() companyController {
	companyService := service.NewCompanyService(db.Conn)
	return companyController{companyService: &companyService}
}

func (c companyController) SaveCompany(ctx *gin.Context) {
	var companyDto dto.CompanyDTO

	if err := ctx.BindJSON(&companyDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	company, err := c.companyService.SaveCompany(companyDto)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	httpresponse.Created(ctx, company)
}

func (c companyController) UpdateProduct(ctx *gin.Context) {

	var productDto dto.CompanyDTO

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}

	if err := ctx.BindJSON(&productDto); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	product, err := c.companyService.UpdateCompany(productDto, id)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	httpresponse.Ok(ctx, product)
}

func (c companyController) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}

	if !c.companyService.DeleteCompany(id) {
		exceptions.MotionError("cannot be deleted").Throw(ctx)
		return
	}
	httpresponse.Ok(ctx, nil)
}

func (c companyController) GetCompany(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	product, err := c.companyService.GetCompany(id)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	httpresponse.Ok(ctx, product)
}

func (c companyController) GetCompanyInfo(ctx *gin.Context) {
	companyName := c.companyService.FindByCompanyName(ctx.Param("companyName"))
	httpresponse.Ok(ctx, companyName)
}
