package repository

import (
	"github.com/ribeirosaimon/motion-go/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	userRepository         *MotionSQLRepository[sqlDomain.MotionUser]
	sessionRepository      *MotionSQLRepository[sqlDomain.Session]
	profileRepository      *MotionSQLRepository[sqlDomain.Profile]
	roleRepository         *MotionSQLRepository[sqlDomain.Role]
	shoppingCartRepository *motionNoSQLRepository[nosqlDomain.ShoppingCart]
	productRepository      *MotionSQLRepository[sqlDomain.Product]
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

func NewShoppingCartRepository(mongoConnection *mongo.Client) *motionNoSQLRepository[nosqlDomain.ShoppingCart] {
	if shoppingCartRepository != nil {
		return shoppingCartRepository
	}
	shoppingCartRepository := newMotionNoSQLRepository[nosqlDomain.ShoppingCart](mongoConnection)
	return shoppingCartRepository
}

func NewProductRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain.Product] {
	if productRepository != nil {
		return productRepository
	}
	productRepository := newMotionSQLRepository[sqlDomain.Product](conn)
	return productRepository
}
