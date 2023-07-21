package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/service"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/httpresponse"
)

type companyController struct {
	companyService *service.CompanyService
}

func NewCompanyController() *companyController {
	companyService := service.NewCompanyService(db.Conn)
	return &companyController{companyService: companyService}
}

func (c *companyController) DeleteProduct(ctx *gin.Context) {
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

func (c *companyController) GetCompany(ctx *gin.Context) {
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

func (c *companyController) GetCompanyInfo(ctx *gin.Context) {
	companyName, err := c.companyService.FindByCompanyCode(ctx.Param("companyName"))
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	httpresponse.Ok(ctx, companyName)
}
