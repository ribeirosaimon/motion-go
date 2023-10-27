package service

import (
	"fmt"
	sqlDomain2 "github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"
	"time"

	"github.com/google/uuid"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type SessionService struct {
	sessionRepository repository.MotionRepository[sqlDomain2.Session]
	roleRepository    repository.MotionRepository[sqlDomain2.Role]
}

func NewSessionService(conn *db.Connections) *SessionService {
	return &SessionService{
		sessionRepository: repository.NewSessionRepository(conn.GetPgsqTemplate()),
		roleRepository:    repository.NewRoleRepository(conn.GetPgsqTemplate()),
	}
}

func (s *SessionService) SaveUserSession(user sqlDomain2.Profile) (sqlDomain2.Session, error) {
	var session sqlDomain2.Session

	session.Id = fmt.Sprintf("%s-%s", uuid.New(), uuid.New())
	session.ProfileId = user.Id
	session.LastLogin = time.Now()

	return s.sessionRepository.Save(session)

}
