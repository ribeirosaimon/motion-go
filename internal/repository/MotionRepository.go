package repository

import (
	"context"
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	userRepository         *MotionSQLRepository[sqlDomain.MotionUser]
	sessionRepository      *MotionSQLRepository[sqlDomain.Session]
	profileRepository      *MotionSQLRepository[sqlDomain.Profile]
	roleRepository         *MotionSQLRepository[sqlDomain.Role]
	shoppingCartRepository *motionNoSQLRepository[nosqlDomain.ShoppingCart]
	companyRepository      *MotionSQLRepository[sqlDomain.Company]
)

func NewUserRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.MotionUser] {
	if userRepository != nil {
		return userRepository
	}
	userRepository := newMotionSQLRepository[sqlDomain.MotionUser](conn)
	return userRepository
}

func NewSessionRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Session] {
	if sessionRepository != nil {
		return sessionRepository
	}
	sessionRepository := newMotionSQLRepository[sqlDomain.Session](conn)
	return sessionRepository
}

func NewProfileRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Profile] {
	if profileRepository != nil {
		return profileRepository
	}
	profileRepository := newMotionSQLRepository[sqlDomain.Profile](conn)
	return profileRepository
}

func NewRoleRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Role] {
	if roleRepository != nil {
		return roleRepository
	}
	roleRepository := newMotionSQLRepository[sqlDomain.Role](conn)
	return roleRepository
}

func NewShoppingCartRepository(ctx context.Context, mongoConnection *mongo.Client) *motionNoSQLRepository[nosqlDomain.ShoppingCart] {
	if shoppingCartRepository != nil {
		return shoppingCartRepository
	}
	shoppingCartRepository := newMotionNoSQLRepository[nosqlDomain.ShoppingCart](ctx, mongoConnection)
	return shoppingCartRepository
}

func NewCompanyRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Company] {
	if companyRepository != nil {
		return companyRepository
	}
	companyRepository := newMotionSQLRepository[sqlDomain.Company](conn)
	return companyRepository
}
