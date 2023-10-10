package service

import (
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type TransactionService struct {
	transactionService repository.MotionRepository[sqlDomain.Transaction]
}

func NewTransactionService(conections *db.Connections) *TransactionService {
	return &TransactionService{
		transactionService: repository.NewTransactionRepository(conections.GetPgsqTemplate()),
	}
}
