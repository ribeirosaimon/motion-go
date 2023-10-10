package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
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

func (s SessionService) SaveUserSession(user sqlDomain.Profile) (sqlDomain.Session, error) {
	var session sqlDomain.Session

	session.SessionId = fmt.Sprintf("%s-%s", uuid.New(), uuid.New())
	session.ProfileId = user.Id
	session.LastLogin = time.Now()

	return s.sessionRepository.Save(session)

}
