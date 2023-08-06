package service

import (
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type TransactionService struct {
	sessionRepository repository.MotionRepository[sqlDomain.Session]
	roleRepository    repository.MotionRepository[sqlDomain.Role]
}
