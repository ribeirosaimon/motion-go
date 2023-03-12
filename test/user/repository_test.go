package user

import (
	"testing"

	"github.com/ribeirosaimon/motion-go/repository"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var userRepository = repository.NewUserRepository(util.ConnectDatabaseTest())

func BenchmarkRepository(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := util.CreateUser()
		user, err := userRepository.Save(user)
		if err != nil {
			b.Errorf("Erro saving %s", user.Name)
		}
	}
}

func TestSaveInRepository(t *testing.T) {
	user := util.CreateUser()
	userSaved, err := userRepository.Save(user)
	if err != nil {
		t.Errorf("Expected to save MotionUser but return error")
	}
	userFound, err := userRepository.FindById(userSaved.Id)
	if err != nil {
		t.Errorf("Expected to save MotionUser but return error")
	}
	if userSaved.Name != userFound.Name {
		t.Errorf("Name must be %s but is %s", userSaved.Name, userFound.Name)
	}
	if userSaved.LoginCount != userFound.LoginCount {
		t.Errorf("NAme must be %d but is %d", userSaved.LoginCount, userFound.LoginCount)
	}
}

func TestUpdateInRepository(t *testing.T) {
	user := util.CreateUser()
	userSaved, err := userRepository.Save(user)
	if err != nil {
		t.Errorf("Expected to save MotionUser but return error")
	}
	userSaved.Name = "newName"
	updateUser, err := userRepository.Save(userSaved)
	userFound, err := userRepository.FindById(updateUser.Id)
	if err != nil {
		t.Errorf("Expected to save MotionUser but return error")
	}
	if userSaved.Name != userFound.Name {
		t.Errorf("Name must be %s but is %s", userSaved.Name, userFound.Name)
	}
	if userSaved.LoginCount != userFound.LoginCount {
		t.Errorf("Name must be %d but is %d", userSaved.LoginCount, userFound.LoginCount)
	}
}

func TestByIdInRepository(t *testing.T) {
	user := util.CreateUser()
	userSaved, err := userRepository.Save(user)
	if err != nil {
		t.Errorf("Expected to save MotionUser but return error")
	}
	userSaved.Name = "newName"
	updateUser, err := userRepository.Save(userSaved)
	userFound, err := userRepository.FindById(updateUser.Id)
	if err != nil {
		t.Errorf("Expected to save MotionUser but return error")
	}
	if userSaved.Name != userFound.Name {
		t.Errorf("NAme must be %s but is %s", userSaved.Name, userFound.Name)
	}
	if userSaved.LoginCount != userFound.LoginCount {
		t.Errorf("NAme must be %d but is %d", userSaved.LoginCount, userFound.LoginCount)
	}
}

func TestDeleteInRepository(t *testing.T) {
	userSaved, err := userRepository.Save(util.CreateUser())
	userRepository.Save(userSaved)
	all, err := userRepository.FindAll(10, 0)
	if len(all) == 0 {
		t.Errorf("Erro saving user")
	}
	user := all[0]
	err = userRepository.DeleteById(user.Id)
	if err != nil {
		t.Errorf("Expected to save MotionUser but return error")
	}
	_, err = userRepository.FindById(user.Id)
	if err == nil {
		t.Errorf("User should not be founded")
	}

}
