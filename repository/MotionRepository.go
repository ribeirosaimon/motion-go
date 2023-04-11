package repository

import (
	"github.com/ribeirosaimon/motion-go/domain"
	"gorm.io/gorm"
)

var (
	userRepository         *MotionRepository[domain.MotionUser]
	sessionRepository      *MotionRepository[domain.Session]
	profileRepository      *MotionRepository[domain.Profile]
	roleRepository         *MotionRepository[domain.Role]
	shoppingCartRepository *MotionRepository[domain.ShoppingCart]
	productRepository      *MotionRepository[domain.Product]
)

func NewUserRepository(conn *gorm.DB) MotionRepository[domain.MotionUser] {
	if userRepository != nil {
		return *userRepository
	}
	return newMotionRepository[domain.MotionUser](conn)
}

func NewSessionRepository(conn *gorm.DB) MotionRepository[domain.Session] {
	if sessionRepository != nil {
		return *sessionRepository
	}
	return newMotionRepository[domain.Session](conn)
}

func NewProfileRepository(conn *gorm.DB) MotionRepository[domain.Profile] {
	if profileRepository != nil {
		return *profileRepository
	}
	return newMotionRepository[domain.Profile](conn)
}

func NewRoleRepository(conn *gorm.DB) MotionRepository[domain.Role] {
	if roleRepository != nil {
		return *roleRepository
	}
	return newMotionRepository[domain.Role](conn)
}

func NewShoppingCartRepository(conn *gorm.DB) MotionRepository[domain.ShoppingCart] {
	if shoppingCartRepository != nil {
		return *shoppingCartRepository
	}
	return newMotionRepository[domain.ShoppingCart](conn)
}

func NewProductRepository(conn *gorm.DB) MotionRepository[domain.Product] {
	if productRepository != nil {
		return *productRepository
	}
	return newMotionRepository[domain.Product](conn)
}
