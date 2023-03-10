package session

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/database"
	"github.com/ribeirosaimon/motion-go/repository"
)

type sessionService struct {
	sessionRepository repository.MotionRepository[domain.Session]
}

func (s sessionService) GetLoggedUserSession(context *gin.Context) {

}

func NewLoginService() sessionService {
	sessionRepository := repository.NewSessionRepository(database.Connect())
	return sessionService{sessionRepository: sessionRepository}
}
