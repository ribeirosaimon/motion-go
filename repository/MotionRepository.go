package repository

import (
	"github.com/ribeirosaimon/motion-go/domain/nosql"
	"github.com/ribeirosaimon/motion-go/domain/sql"
	"gorm.io/gorm"
)

var (
	userRepository         *MotionRepository[sql.MotionUser]
	sessionRepository      *MotionRepository[sql.Session]
	profileRepository      *MotionRepository[sql.Profile]
	roleRepository         *MotionRepository[sql.Role]
	shoppingCartRepository *MotionRepository[nosql.ShoppingCart]
	productRepository      *MotionRepository[sql.Product]
)

func NewUserRepository(conn *gorm.DB) MotionRepository[sql.MotionUser] {
	if userRepository != nil {
		return *userRepository
	}
	return newMotionRepository[sql.MotionUser](conn)
}

func NewSessionRepository(conn *gorm.DB) MotionRepository[sql.Session] {
	if sessionRepository != nil {
		return *sessionRepository
	}
	return newMotionRepository[sql.Session](conn)
}

func NewProfileRepository(conn *gorm.DB) MotionRepository[sql.Profile] {
	if profileRepository != nil {
		return *profileRepository
	}
	return newMotionRepository[sql.Profile](conn)
}

func NewRoleRepository(conn *gorm.DB) MotionRepository[sql.Role] {
	if roleRepository != nil {
		return *roleRepository
	}
	return newMotionRepository[sql.Role](conn)
}

func NewShoppingCartRepository(conn *gorm.DB) MotionRepository[nosql.ShoppingCart] {
	if shoppingCartRepository != nil {
		return *shoppingCartRepository
	}
	return newMotionRepository[nosql.ShoppingCart](conn)
}

func NewProductRepository(conn *gorm.DB) MotionRepository[sql.Product] {
	if productRepository != nil {
		return *productRepository
	}
	return newMotionRepository[sql.Product](conn)
}
