package session

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	repository2 "github.com/ribeirosaimon/motion-go/internal/repository"
	"gorm.io/gorm"
)

type Service struct {
	sessionRepository repository2.MotionRepository[sqlDomain.Session]
	roleRepository    repository2.MotionRepository[sqlDomain.Role]
	closeDb           *sql.DB
}

func NewLoginService(conn *gorm.DB, close *sql.DB) Service {
	return Service{sessionRepository: repository2.NewSessionRepository(conn),
		roleRepository: repository2.NewRoleRepository(conn),
		closeDb:        close,
	}
}

func (s Service) SaveUserSession(user sqlDomain.Profile) (sqlDomain.Session, error) {
	var session sqlDomain.Session

	session.SessionId = fmt.Sprintf("%s-%s", uuid.New(), uuid.New())
	session.ProfileId = user.Id
	session.LastLogin = time.Now()

	return s.sessionRepository.Save(session)

}
