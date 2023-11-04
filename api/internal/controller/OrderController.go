package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/dto"
	"github.com/ribeirosaimon/motion-go/api/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/api/internal/middleware"
	"github.com/ribeirosaimon/motion-go/api/internal/response"
	"github.com/ribeirosaimon/motion-go/api/internal/service"
)

type orderController struct {
	orderService *service.OrderService
}

func NewOrderController() *orderController {
	orderService := service.NewOrderService(db.Conn)
	return &orderController{orderService: orderService}
}

func (c *orderController) NewOrder(ctx *gin.Context) {
	loggedUser := middleware.GetLoggedUser(ctx)
	var order dto.Order
	if err := ctx.BindJSON(&order); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}

	if err := c.orderService.NewOrder(loggedUser, order); err != nil {
		exceptions.InternalServer("message").Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusCreated, nil)
}

func (c *orderController) FindAll(ctx *gin.Context) {
	loggedUser := middleware.GetLoggedUser(ctx)
	allOrder, err := c.orderService.FindAll(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusOK, allOrder)
}
