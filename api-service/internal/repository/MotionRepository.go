package repository

import (
	"context"
	nosqlDomain2 "github.com/ribeirosaimon/motion-go/config/domain/nosqlDomain"
	sqlDomain2 "github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	userRepository         *MotionSQLRepository[sqlDomain2.MotionUser]
	sessionRepository      *MotionSQLRepository[sqlDomain2.Session]
	profileRepository      *MotionSQLRepository[sqlDomain2.Profile]
	roleRepository         *MotionSQLRepository[sqlDomain2.Role]
	transactionRepository  *MotionSQLRepository[sqlDomain2.Transaction]
	watchListRepository    *MotionNoSQLRepository[nosqlDomain2.WatchList]
	summaryStockRepository *MotionNoSQLRepository[nosqlDomain2.SummaryStock]
)

func NewUserRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain2.MotionUser] {
	if userRepository != nil {
		return userRepository
	}
	userRepository := newMotionSQLRepository[sqlDomain2.MotionUser](conn)
	return userRepository
}

func NewSessionRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain2.Session] {
	if sessionRepository != nil {
		return sessionRepository
	}
	sessionRepository := newMotionSQLRepository[sqlDomain2.Session](conn)
	return sessionRepository
}

func NewProfileRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain2.Profile] {
	if profileRepository != nil {
		return profileRepository
	}
	profileRepository := newMotionSQLRepository[sqlDomain2.Profile](conn)
	return profileRepository
}

func NewRoleRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain2.Role] {
	if roleRepository != nil {
		return roleRepository
	}
	roleRepository := newMotionSQLRepository[sqlDomain2.Role](conn)
	return roleRepository
}

func NewWatchListRepository(ctx context.Context, mongoConnection *mongo.Client) *MotionNoSQLRepository[nosqlDomain2.WatchList] {
	if watchListRepository != nil {
		return watchListRepository
	}
	watchListRepository := newMotionNoSQLRepository[nosqlDomain2.WatchList](ctx, mongoConnection)
	return watchListRepository
}

func NewSummaryStockRepository(ctx context.Context, mongoConnection *mongo.Client) *MotionNoSQLRepository[nosqlDomain2.SummaryStock] {
	if summaryStockRepository != nil {
		return summaryStockRepository
	}
	summaryStockRepository := newMotionNoSQLRepository[nosqlDomain2.SummaryStock](ctx, mongoConnection)
	return summaryStockRepository
}

func NewTransactionRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain2.Transaction] {
	if transactionRepository != nil {
		return transactionRepository
	}
	transactionRepository := newMotionSQLRepository[sqlDomain2.Transaction](conn)
	return transactionRepository
}
