package user

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/repository"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var userRepository = repository.NewUserRepository(util.ConnectDatabaseTest())

func createUser() domain.MotionUser {
	rand.Seed(time.Now().UnixNano())
	nameRandom := strconv.Itoa(rand.Intn(1000000))
	lastNameRandom := strconv.Itoa(rand.Intn(1000000))
	emailRandom := fmt.Sprintf("%s@email.com", strconv.Itoa(rand.Intn(1000000)))
	LoginAttempRandom := uint8(rand.Intn(10))
	loginCountRandom := uint64(rand.Intn(101))

	return domain.MotionUser{
		Name:        nameRandom,
		LastName:    lastNameRandom,
		LoginCount:  loginCountRandom,
		Email:       emailRandom,
		LoginAttemp: LoginAttempRandom,
	}

}
func BenchmarkRepository(b *testing.B) {

}

func TestSaveInRepository(t *testing.T) {
	user := createUser()
	userSaved, err := userRepository.Save(user)
	if err != nil {
		t.Errorf("Expected to save MotionUser but return error")
	}
	userFound, err := userRepository.FindById(userSaved.Id)
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

func TestUpdateInRepository(t *testing.T) {
	user := createUser()
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

func TestFindAllInRepository(t *testing.T) {
	user := createUser()
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
