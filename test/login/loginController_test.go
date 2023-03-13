package login

import (
	"fmt"
	"testing"

	"github.com/ribeirosaimon/motion-go/repository"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var userRepository = repository.NewUserRepository(util.ConnectDatabaseTest())

func TestLoginController(t *testing.T) {
	user := util.SaveUserInDb()
	id, err := userRepository.FindById(user.Id)
	sid, err := userRepository.FindAll(10, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(sid)
	session := util.GetUserSession(id)
	fmt.Println(session)
}
