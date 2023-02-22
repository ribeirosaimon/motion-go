package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/repository"
)

func getUserService(ctx *gin.Context) {
	userRepository := repository.NewUserRepository()
	userRepository.FindById("teste")
}
