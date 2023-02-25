package repository

import "github.com/ribeirosaimon/motion-go/domain"

func NewUserRepository() motionRepository[domain.MotionUser] {
	return newMotionRepository[domain.MotionUser]()
}
