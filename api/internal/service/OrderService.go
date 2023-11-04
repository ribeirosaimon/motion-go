package service

import (
	"errors"

	"github.com/ribeirosaimon/motion-go/api/internal/akafka"
	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/dto"
	"github.com/ribeirosaimon/motion-go/api/internal/middleware"
	"github.com/ribeirosaimon/motion-go/api/internal/repository"
	"github.com/ribeirosaimon/motion-go/confighub/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/confighub/domain/sqlDomain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	orderRepository        *repository.MotionSQLRepository[sqlDomain.Order]
	summaryStockRepository *repository.MotionNoSQLRepository[nosqlDomain.SummaryStock]
	motionKafka            *akafka.MotionKafka
}

func NewOrderService(conn *db.Connections) *OrderService {
	return &OrderService{
		orderRepository:        repository.NewOrderRepository(conn.GetPgsqTemplate()),
		summaryStockRepository: repository.NewSummaryStockRepository(conn.Context, conn.GetMongoTemplate()),
		motionKafka:            akafka.GetMotionKafka(),
	}
}

func (s *OrderService) NewOrder(user middleware.LoggedUser, order dto.Order) error {

	objectID, err := primitive.ObjectIDFromHex(order.Id)
	if err != nil {
		return err
	}
	stock, err := s.summaryStockRepository.FindByField("_id", objectID)

	if err != nil {
		return errors.New("this company does not exist")
	}

	newOrder := sqlDomain.NewOrder(order.Value, order.Quantity, user.ProfileId, user.SessionId, objectID.Hex(), stock.CompanyCode, sqlDomain.OrderActive)

	savedOrder, err := s.orderRepository.Save(newOrder)
	if err != nil {
		return err
	}
	if err = s.motionKafka.SendMessage(savedOrder); err != nil {
		return err
	}
	return nil
}

func (s *OrderService) FindAll(loggedUser middleware.LoggedUser) ([]sqlDomain.Order, error) {
	if loggedUser.Role.Name == sqlDomain.ADMIN {
		allTransactions, err := s.orderRepository.FindAll(10, 0)
		if err != nil {
			return []sqlDomain.Order{}, err
		}
		return allTransactions, nil
	}
	var query = "SELECT o.* FROM orders o WHERE o.profile_id = ?"
	transactionRepository := repository.NewOrderRepository(db.Conn.GetPgsqTemplate())
	transactions := []sqlDomain.Order{}
	originalSlice, err := transactionRepository.CreateNativeSQLQuery(query, transactions, loggedUser.ProfileId)
	if err != nil {
		return []sqlDomain.Order{}, err
	}

	return originalSlice.([]sqlDomain.Order), nil
}
