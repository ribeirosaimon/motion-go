package repository

import (
	"context"

	"github.com/ribeirosaimon/motion-go/confighub/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/confighub/domain/sqlDomain"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	userRepository         *MotionSQLRepository[sqlDomain.MotionUser]
	sessionRepository      *MotionSQLRepository[sqlDomain.Session]
	profileRepository      *MotionSQLRepository[sqlDomain.Profile]
	roleRepository         *MotionSQLRepository[sqlDomain.Role]
	transactionRepository  *MotionSQLRepository[sqlDomain.Transaction]
	orderRepository        *MotionSQLRepository[sqlDomain.Order]
	watchListRepository    *MotionNoSQLRepository[nosqlDomain.WatchList]
	summaryStockRepository *MotionNoSQLRepository[nosqlDomain.SummaryStock]
)

func NewUserRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.MotionUser] {
	if userRepository == nil {
		userRepository = newMotionSQLRepository[sqlDomain.MotionUser](conn)
	}
	return userRepository
}

func NewSessionRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Session] {
	if sessionRepository == nil {
		sessionRepository = newMotionSQLRepository[sqlDomain.Session](conn)
	}
	return sessionRepository
}

func NewProfileRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Profile] {
	if profileRepository == nil {
		profileRepository = newMotionSQLRepository[sqlDomain.Profile](conn)
	}
	return profileRepository
}

func NewRoleRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Role] {
	if roleRepository == nil {
		roleRepository = newMotionSQLRepository[sqlDomain.Role](conn)
	}
	return roleRepository
}

func NewWatchListRepository(ctx context.Context, mongoConnection *mongo.Client) *MotionNoSQLRepository[nosqlDomain.WatchList] {
	if watchListRepository == nil {
		watchListRepository = newMotionNoSQLRepository[nosqlDomain.WatchList](ctx, mongoConnection)
	}
	return watchListRepository
}

func NewSummaryStockRepository(ctx context.Context, mongoConnection *mongo.Client) *MotionNoSQLRepository[nosqlDomain.SummaryStock] {
	if summaryStockRepository == nil {
		summaryStockRepository = newMotionNoSQLRepository[nosqlDomain.SummaryStock](ctx, mongoConnection)
	}
	return summaryStockRepository
}

func NewTransactionRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Transaction] {
	if transactionRepository == nil {
		transactionRepository = newMotionSQLRepository[sqlDomain.Transaction](conn)
	}
	return transactionRepository
}

func NewOrderRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Order] {
	if orderRepository == nil {
		orderRepository = newMotionSQLRepository[sqlDomain.Order](conn)
	}
	return orderRepository
}
