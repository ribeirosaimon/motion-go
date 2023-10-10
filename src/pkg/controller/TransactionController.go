package controller

import (
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/src/pkg/service"
)

type transactionController struct {
	transactionService *service.TransactionService
}

func NewTransactionController() *transactionController {
	ts := service.NewTransactionService(db.Conn)

	return &transactionController{transactionService: ts}
}
