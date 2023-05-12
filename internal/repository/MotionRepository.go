package repository

import (
	"github.com/ribeirosaimon/motion-go/internal/domain/nosqlDomain"
	sqlDomain2 "github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	userRepository         *MotionSQLRepository[sqlDomain2.MotionUser]
	sessionRepository      *MotionSQLRepository[sqlDomain2.Session]
	profileRepository      *MotionSQLRepository[sqlDomain2.Profile]
	roleRepository         *MotionSQLRepository[sqlDomain2.Role]
	shoppingCartRepository *motionNoSQLRepository[nosqlDomain.ShoppingCart]
	productRepository      *MotionSQLRepository[sqlDomain2.Product]
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

func NewShoppingCartRepository(mongoConnection *mongo.Client) *motionNoSQLRepository[nosqlDomain.ShoppingCart] {
	if shoppingCartRepository != nil {
		return shoppingCartRepository
	}
	shoppingCartRepository := newMotionNoSQLRepository[nosqlDomain.ShoppingCart](mongoConnection)
	return shoppingCartRepository
}

func NewProductRepository(conn *gorm.DB) *MotionSQLRepository[sqlDomain2.Product] {
	if productRepository != nil {
		return productRepository
	}
	productRepository := newMotionSQLRepository[sqlDomain2.Product](conn)
	return productRepository
}
