package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/repository"
	"time"
)

func getUserService(ctx *gin.Context) {
	userRepository := repository.NewUserRepository()

	user := domain.User{
		Name:        "Saimon",
		LoginAttemp: 0,
		LastLogin:   time.Now(),
	}

	userRepository.Save(user)
}
