package repository

import (
	"github.com/ribeirosaimon/motion-go/domain"
	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) MotionRepository[domain.MotionUser] {
	return newMotionRepository[domain.MotionUser](conn)
}
