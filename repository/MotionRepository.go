package repository

import (
	"github.com/ribeirosaimon/motion-go/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"gorm.io/gorm"
)

var (
	userRepository         *motionSQLRepository[sqlDomain.MotionUser]
	sessionRepository      *motionSQLRepository[sqlDomain.Session]
	profileRepository      *motionSQLRepository[sqlDomain.Profile]
	roleRepository         *motionSQLRepository[sqlDomain.Role]
	shoppingCartRepository *motionNoSQLRepository[nosqlDomain.ShoppingCart]
	productRepository      *motionSQLRepository[sqlDomain.Product]
)

func NewUserRepository(conn *gorm.DB) *motionSQLRepository[sqlDomain.MotionUser] {
	if userRepository != nil {
		return userRepository
	}
	userRepository := newMotionSQLRepository[sqlDomain.MotionUser](conn)
	return userRepository
}

func NewSessionRepository(conn *gorm.DB) *motionSQLRepository[sqlDomain.Session] {
	if sessionRepository != nil {
		return sessionRepository
	}
	sessionRepository := newMotionSQLRepository[sqlDomain.Session](conn)
	return sessionRepository
}

func NewProfileRepository(conn *gorm.DB) *motionSQLRepository[sqlDomain.Profile] {
	if profileRepository != nil {
		return profileRepository
	}
	profileRepository := newMotionSQLRepository[sqlDomain.Profile](conn)
	return profileRepository
}

func NewRoleRepository(conn *gorm.DB) *motionSQLRepository[sqlDomain.Role] {
	if roleRepository != nil {
		return roleRepository
	}
	roleRepository := newMotionSQLRepository[sqlDomain.Role](conn)
	return roleRepository
}

func NewShoppingCartRepository(conn *gorm.DB) *motionNoSQLRepository[nosqlDomain.ShoppingCart] {
	if shoppingCartRepository != nil {
		return shoppingCartRepository
	}
	shoppingCartRepository := newMotionNoSQLRepository[nosqlDomain.ShoppingCart](conn)
	return shoppingCartRepository
}

func NewProductRepository(conn *gorm.DB) motionSQLRepository[sqlDomain.Product] {
	if productRepository != nil {
		return *productRepository
	}
	return newMotionSQLRepository[sqlDomain.Product](conn)
}
