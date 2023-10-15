package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/dto"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/response"
	"github.com/ribeirosaimon/motion-go/internal/service"
)

type TransactionController struct {
	transactionService *service.TransactionService
}

func NewTransactionController() *TransactionController {
	ts := service.NewTransactionService(db.Conn)

	return &TransactionController{transactionService: ts}
}

func (t *TransactionController) Deposit(ctx *gin.Context) {
	loggedUser, err := middleware.GetLoggedUser(ctx)
	var deposit dto.Deposit
	if err := ctx.BindJSON(&deposit); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	transaction, err := t.transactionService.Deposit(loggedUser, deposit)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusCreated, transaction)
}
