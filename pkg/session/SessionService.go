package session

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	sql2 "github.com/ribeirosaimon/motion-go/domain/sql"
	"github.com/ribeirosaimon/motion-go/repository"
	"gorm.io/gorm"
)

type Service struct {
	sessionRepository repository.MotionRepository[sql2.Session]
	roleRepository    repository.MotionRepository[sql2.Role]
	closeDb           *sql.DB
}

func NewLoginService(conn *gorm.DB, close *sql.DB) Service {
	return Service{sessionRepository: repository.NewSessionRepository(conn),
		roleRepository: repository.NewRoleRepository(conn),
		closeDb:        close,
	}
}

func (s Service) SaveUserSession(user sql2.Profile) (sql2.Session, error) {
	var session sql2.Session

	session.SessionId = fmt.Sprintf("%s-%s", uuid.New(), uuid.New())
	session.ProfileId = user.Id
	session.LastLogin = time.Now()

	return s.sessionRepository.Save(session)

}
