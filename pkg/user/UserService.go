package user

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/repository"
)

func getUserService(ctx *gin.Context) {
	userRepository := repository.NewUserRepository()
	var senha uint64
	senha = 1
	id, err := userRepository.FindById(senha)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Printf("id numero %d e nome %s\n", id.Id, id.Name)
}

func deleteUser(engine *gin.Context) {
	userRepository := repository.NewUserRepository()
	err := userRepository.DeleteById(uint64(2))
	fmt.Errorf(err.Error())
}

func updateUserService(engine *gin.Context) {
	userRepository := repository.NewUserRepository()
	id, err := userRepository.FindById(2)
	if err != nil {
		fmt.Errorf("error")
	}
	id.Name = "Saimon Att"
	save, _ := userRepository.Save(id)
	fmt.Println(save)
}

func saveUserService(engine *gin.Context) {
	userRepository := repository.NewUserRepository()

	user := domain.MotionUser{
		Name:        "Saimon",
		LoginAttemp: 0,
		LastName:    "Ribeiro",
		LastLogin:   time.Now(),
	}

	save, _ := userRepository.Save(user)
	fmt.Println(save)
}
