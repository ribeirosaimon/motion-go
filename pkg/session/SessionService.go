package session

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/repository"
	"gorm.io/gorm"
)

type Service struct {
	sessionRepository repository.MotionRepository[domain.Session]
	roleRepository    repository.MotionRepository[domain.Role]
	closeDb           *sql.DB
}

func NewLoginService(conn *gorm.DB, close *sql.DB) Service {
	return Service{sessionRepository: repository.NewSessionRepository(conn),
		roleRepository: repository.NewRoleRepository(conn),
		closeDb:        close,
	}
}

func (s Service) SaveUserSession(user domain.Profile) (domain.Session, error) {
	var session domain.Session

	session.SessionId = fmt.Sprintf("%s-%s", uuid.New(), uuid.New())
	session.ProfileId = user.Id
	session.LastLogin = time.Now()

	return s.sessionRepository.Save(session)

}
