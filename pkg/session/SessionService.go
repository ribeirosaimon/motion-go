package session

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/database"
	"github.com/ribeirosaimon/motion-go/repository"
)

type Service struct {
	sessionRepository repository.MotionRepository[domain.Session]
	roleRepository    repository.MotionRepository[domain.Role]
}

func NewLoginService() Service {
	connect, _ := database.Connect()
	return Service{sessionRepository: repository.NewSessionRepository(connect),
		roleRepository: repository.NewRoleRepository(connect)}
}

func (s Service) SaveUserSession(user domain.Profile) (domain.Session, error) {
	var session domain.Session

	session.SessionId = fmt.Sprintf("%s-%s", uuid.New(), uuid.New())
	session.ProfileId = user.Id
	session.LastLogin = time.Now()

	return s.sessionRepository.Save(session)

}
