package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/repository"
)

func getUserService(ctx *gin.Context) {
	userRepository := repository.NewUserRepository()

	user := domain.User{
		Name:        "Saimon",
		Id:          123123,
		LoginAttemp: 0,
		LastLogin:   time.Now(),
	}

	userRepository.Save(user)
}
