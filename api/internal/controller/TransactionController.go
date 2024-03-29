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

type TransactionController struct {
	transactionService *service.TransactionService
}

func NewTransactionController() *TransactionController {
	ts := service.NewTransactionService(db.Conn)

	return &TransactionController{transactionService: ts}
}

func (t *TransactionController) Deposit(ctx *gin.Context) {
	loggedUser := middleware.GetLoggedUser(ctx)
	var deposit dto.Deposit
	if err := ctx.BindJSON(&deposit); err != nil {
		exceptions.BodyError().Throw(ctx)
		return
	}

	transaction, err := t.transactionService.Deposit(loggedUser, deposit)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusCreated, transaction)
}

func (t *TransactionController) Balance(ctx *gin.Context) {
	loggedUser := middleware.GetLoggedUser(ctx)
	balance, err := t.transactionService.Balance(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusOK, balance)
}

func (t *TransactionController) FindAllTransactions(ctx *gin.Context) {
	loggedUser := middleware.GetLoggedUser(ctx)
	balance, err := t.transactionService.FindAllTransactions(loggedUser)
	if err != nil {
		exceptions.MotionError(err.Error()).Throw(ctx)
		return
	}
	response.Entity(ctx, http.StatusOK, balance)
}
