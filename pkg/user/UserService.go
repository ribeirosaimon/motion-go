package user

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/database"
	"github.com/ribeirosaimon/motion-go/repository"
)

type userService struct {
	userRepository repository.MotionRepository[domain.MotionUser]
}

func newUserService() userService {
	userRepository := repository.NewUserRepository(database.Connect())
	return userService{
		userRepository: userRepository,
	}
}
func (s userService) getUserService(ctx *gin.Context) {
	// var senha uint64
	// senha = 1
	id, err := s.userRepository.FindAll(10, 0)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(id)
	// fmt.Printf("id numero %d e nome %s\n", id.Id, id.Name)
}

func (s userService) deleteUser(engine *gin.Context) {
	err := s.userRepository.DeleteById(uint64(2))
	fmt.Errorf(err.Error())
}

func (s userService) updateUserService(engine *gin.Context) {

	id, err := s.userRepository.FindById(2)
	if err != nil {
		fmt.Errorf("error")
	}
	id.Name = "Saimon Att"
	save, _ := s.userRepository.Save(id)
	fmt.Println(save)
}

func (s userService) saveUserService(engine *gin.Context) {
	user := domain.MotionUser{
		Name:        "Saimon",
		LoginAttemp: 0,
		LastName:    "Ribeiro",
		LastLogin:   time.Now(),
	}

	save, _ := s.userRepository.Save(user)
	fmt.Println(save)
}
