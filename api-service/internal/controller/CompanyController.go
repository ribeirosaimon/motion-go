package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/response"
	"github.com/ribeirosaimon/motion-go/internal/service"
)

type companyController struct {
	companyService *service.CompanyService
}

func NewCompanyController() *companyController {
	companyService := service.NewCompanyService(db.Conn)
	return &companyController{companyService: companyService}
}

func (c *companyController) DeleteCompany(ctx *gin.Context) {
	id := ctx.Param("id")
	if !c.companyService.DeleteCompany(id) {
		exceptions.MotionError("cannot be deleted").Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusOK, nil)
}

func (c *companyController) GetCompany(ctx *gin.Context) {
	product, err := c.companyService.GetCompany(ctx.Param("id"))
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusOK, product)
}

func (c *companyController) GetCompanyInfo(ctx *gin.Context) {
	companyName, err := c.companyService.FindByCompanyByCodeOrName(ctx.Param("companyCode"), false)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusOK, companyName)
}

func (c *companyController) GetAllCompany(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)
	limit, err := strconv.ParseUint(ctx.Query("limit"), 10, 32)
	allCompany, err := c.companyService.FindAllCompany(uint32(limit), uint32(page))
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusOK, allCompany)
}
