package repository

import (
	"github.com/ribeirosaimon/motion-go/domain/nosqlDomain"
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"gorm.io/gorm"
)

var (
	userRepository         *MotionRepository[sqlDomain.MotionUser]
	sessionRepository      *MotionRepository[sqlDomain.Session]
	profileRepository      *MotionRepository[sqlDomain.Profile]
	roleRepository         *MotionRepository[sqlDomain.Role]
	shoppingCartRepository *MotionRepository[nosqlDomain.ShoppingCart]
	productRepository      *MotionRepository[sqlDomain.Product]
)

func NewUserRepository(conn *gorm.DB) MotionRepository[sqlDomain.MotionUser] {
	if userRepository != nil {
		return *userRepository
	}
	return newMotionSQLRepository[sqlDomain.MotionUser](conn)
}

func NewSessionRepository(conn *gorm.DB) MotionRepository[sqlDomain.Session] {
	if sessionRepository != nil {
		return *sessionRepository
	}
	return newMotionSQLRepository[sqlDomain.Session](conn)
}

func NewProfileRepository(conn *gorm.DB) MotionRepository[sqlDomain.Profile] {
	if profileRepository != nil {
		return *profileRepository
	}
	return newMotionSQLRepository[sqlDomain.Profile](conn)
}

func NewRoleRepository(conn *gorm.DB) MotionRepository[sqlDomain.Role] {
	if roleRepository != nil {
		return *roleRepository
	}
	return newMotionSQLRepository[sqlDomain.Role](conn)
}

func NewShoppingCartRepository(conn *gorm.DB) MotionRepository[nosqlDomain.ShoppingCart] {
	if shoppingCartRepository != nil {
		return *shoppingCartRepository
	}
	return newMotionSQLRepository[nosqlDomain.ShoppingCart](conn)
}

func NewProductRepository(conn *gorm.DB) MotionRepository[sqlDomain.Product] {
	if productRepository != nil {
		return *productRepository
	}
	return newMotionSQLRepository[sqlDomain.Product](conn)
}
