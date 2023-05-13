package session

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type Service struct {
	sessionRepository repository.MotionRepository[sqlDomain.Session]
	roleRepository    repository.MotionRepository[sqlDomain.Role]
	closeDb           *sql.DB
}

func NewLoginService(conn *db.Connections) Service {
	return Service{sessionRepository: repository.NewSessionRepository(conn.sqlStruct.Conn),
		roleRepository: repository.NewRoleRepository(conn.sqlStruct.Conn),
	}
}

func (s Service) SaveUserSession(user sqlDomain.Profile) (sqlDomain.Session, error) {
	var session sqlDomain.Session

	session.SessionId = fmt.Sprintf("%s-%s", uuid.New(), uuid.New())
	session.ProfileId = user.Id
	session.LastLogin = time.Now()

	return s.sessionRepository.Save(session)

}
