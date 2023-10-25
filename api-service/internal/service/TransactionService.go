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

	profile, err := s.profileService.FindProfileByUserId(loggedUser.ProfileId)
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

func (s *TransactionService) FindAllTransactions(loggedUser middleware.LoggedUser) ([]sqlDomain.Transaction, error) {
	if loggedUser.Role.Name == sqlDomain.ADMIN {
		allTransactions, err := s.transactionRepository.FindAll(10, 0)
		if err != nil {
			return []sqlDomain.Transaction{}, err
		}
		return allTransactions, nil
	}
	var query = "SELECT t.* FROM transactions t WHERE t.profile_id = ?"
	transactionRepository := repository.NewTransactionRepository(db.Conn.GetPgsqTemplate())
	transactions := []sqlDomain.Transaction{}
	originalSlice, err := transactionRepository.CreateNativeSQLQuery(query, transactions, loggedUser.ProfileId)
	if err != nil {
		return []sqlDomain.Transaction{}, err
	}

	return originalSlice.([]sqlDomain.Transaction), nil
}