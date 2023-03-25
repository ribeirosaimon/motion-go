package repository

import (
	"github.com/ribeirosaimon/motion-go/domain"
	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) MotionRepository[domain.MotionUser] {
	return newMotionRepository[domain.MotionUser](conn)
}

func NewSessionRepository(conn *gorm.DB) MotionRepository[domain.Session] {
	return newMotionRepository[domain.Session](conn)
}

func NewProfileRepository(conn *gorm.DB) MotionRepository[domain.Profile] {
	return newMotionRepository[domain.Profile](conn)
}

func NewRoleRepository(conn *gorm.DB) MotionRepository[domain.Role] {
	return newMotionRepository[domain.Role](conn)
}

func NewShoppingCartRepository(conn *gorm.DB) MotionRepository[domain.ShoppingCart] {
	return newMotionRepository[domain.ShoppingCart](conn)
}
