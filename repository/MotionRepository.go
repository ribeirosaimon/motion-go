package repository

import "github.com/ribeirosaimon/motion-go/domain"

func NewUserRepository() motionRepository[domain.User] {
	return newMotionRepository[domain.User]()
}
