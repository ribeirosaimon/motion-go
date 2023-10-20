package service

import (
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/dto"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type TransactionService struct {
	transactionRepository repository.MotionRepository[sqlDomain.Transaction]
	sessionRepository     repository.MotionRepository[sqlDomain.Session]
	profileService        ProfileService
}

func NewTransactionService(conections *db.Connections) *TransactionService {
	return &TransactionService{
		transactionRepository: repository.NewTransactionRepository(conections.GetPgsqTemplate()),
		sessionRepository:     repository.NewSessionRepository(conections.GetPgsqTemplate()),
		profileService:        *NewProfileService(db.Conn),
	}
}

func (s *TransactionService) Deposit(loggedUser middleware.LoggedUser, deposit dto.Deposit) (sqlDomain.Transaction, error) {

	var transaction sqlDomain.Transaction
	session, err := s.sessionRepository.FindByField("id", loggedUser.SessionId)
	if err != nil {
		return transaction, err
	}

	profile, err := s.profileService.FindProfileByUserId(loggedUser.UserId)
	if err != nil {
		return transaction, err
	}

	transaction.OperationType = sqlDomain.DEPOSIT
	transaction.Value = deposit.Value
	transaction.SessionId = session.Id
	transaction.ProfileId = profile.Id
	var today = time.Now()
	transaction.CreatedAt = today
	transaction.UpdatedAt = today
	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return sqlDomain.Transaction{}, err
	}

	return transaction, nil
}

func (s *TransactionService) Balance(loggedUser middleware.LoggedUser) error {

	// var transaction sqlDomain.Transaction
	// session, err := s.sessionRepository.FindById(loggedUser.SessionId)
	// profile, err := s.profileService.FindProfileByUserId(loggedUser.UserId)
	// if err != nil {
	// 	return err
	// }
	//
	// var query = "SELECT SUM() FROM "

	return nil
}
