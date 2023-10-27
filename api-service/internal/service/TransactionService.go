package service

import (
	sqlDomain2 "github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/dto"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type TransactionService struct {
	transactionRepository repository.MotionRepository[sqlDomain2.Transaction]
	sessionRepository     repository.MotionRepository[sqlDomain2.Session]
	profileService        ProfileService
}

func NewTransactionService(conections *db.Connections) *TransactionService {

	return &TransactionService{
		transactionRepository: repository.NewTransactionRepository(conections.GetPgsqTemplate()),
		sessionRepository:     repository.NewSessionRepository(conections.GetPgsqTemplate()),
		profileService:        *NewProfileService(db.Conn),
	}
}

func (s *TransactionService) Deposit(loggedUser middleware.LoggedUser, deposit dto.Deposit) (sqlDomain2.Transaction, error) {

	var transaction sqlDomain2.Transaction
	session, err := s.sessionRepository.FindByField("id", loggedUser.SessionId)
	if err != nil {
		return transaction, err
	}

	profile, err := s.profileService.FindProfileByUserId(loggedUser.ProfileId)
	if err != nil {
		return transaction, err
	}

	transaction.OperationType = sqlDomain2.DEPOSIT
	transaction.Value = deposit.Value
	transaction.SessionId = session.Id
	transaction.ProfileId = profile.Id
	var today = time.Now()
	transaction.CreatedAt = today
	transaction.UpdatedAt = today
	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return sqlDomain2.Transaction{}, err
	}

	return transaction, nil
}

func (s *TransactionService) Balance(loggedUser middleware.LoggedUser) (dto.Deposit, error) {
	profile, err := s.profileService.FindProfileByUserId(loggedUser.ProfileId)
	var value *dto.Deposit
	if err != nil {
		return *value, err
	}

	var query = "SELECT COALESCE(SUM(t.value),0) FROM transactions t WHERE t.profile_id = ?"

	var total float64
	_, err = repository.NewTransactionRepository(db.Conn.GetPgsqTemplate()).CreateNativeSQLQuery(query, &total, profile.Id)
	if err != nil {
		return dto.Deposit{}, err
	}

	return dto.Deposit{Value: total}, nil
}

func (s *TransactionService) FindAllTransactions(loggedUser middleware.LoggedUser) ([]sqlDomain2.Transaction, error) {
	if loggedUser.Role.Name == sqlDomain2.ADMIN {
		allTransactions, err := s.transactionRepository.FindAll(10, 0)
		if err != nil {
			return []sqlDomain2.Transaction{}, err
		}
		return allTransactions, nil
	}
	var query = "SELECT t.* FROM transactions t WHERE t.profile_id = ?"
	transactionRepository := repository.NewTransactionRepository(db.Conn.GetPgsqTemplate())
	transactions := []sqlDomain2.Transaction{}
	originalSlice, err := transactionRepository.CreateNativeSQLQuery(query, transactions, loggedUser.ProfileId)
	if err != nil {
		return []sqlDomain2.Transaction{}, err
	}

	return originalSlice.([]sqlDomain2.Transaction), nil
}
