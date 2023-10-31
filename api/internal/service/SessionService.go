package service

import (
	"fmt"
	"time"

	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/repository"
	"github.com/ribeirosaimon/motion-go/confighub/domain/sqlDomain"

	"github.com/google/uuid"
)

type SessionService struct {
	sessionRepository repository.MotionRepository[sqlDomain.Session]
	roleRepository    repository.MotionRepository[sqlDomain.Role]
}

func NewSessionService(conn *db.Connections) *SessionService {
	return &SessionService{
		sessionRepository: repository.NewSessionRepository(conn.GetPgsqTemplate()),
		roleRepository:    repository.NewRoleRepository(conn.GetPgsqTemplate()),
	}
}

func (s *SessionService) SaveUserSession(user sqlDomain.Profile) (sqlDomain.Session, error) {
	var session sqlDomain.Session

	session.Id = fmt.Sprintf("%s-%s", uuid.New(), uuid.New())
	session.ProfileId = user.Id
	session.LastLogin = time.Now()

	return s.sessionRepository.Save(session)

}
