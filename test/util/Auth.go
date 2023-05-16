package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/login"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

func SignUp(loggedRole sqlDomain.RoleEnum, roles ...sqlDomain.RoleEnum) (string, error) {
	user := CreateUser(roles...)
	fmt.Sprintf("%s", user)
	// jsonData, err := json.Marshal(user)
	return "", nil
}

func CreateUser(roles ...sqlDomain.RoleEnum) login.SignUpDto {
	createRoles()
	rand.Seed(time.Now().UnixNano())
	nameRandom := strconv.Itoa(rand.Intn(1000000))
	password := strconv.Itoa(rand.Intn(1000000))
	emailRandom := fmt.Sprintf("%s@email.com", strconv.Itoa(rand.Intn(1000000)))

	var dto login.SignUpDto
	dto.Name = nameRandom
	dto.Password = password
	dto.Email = emailRandom
	dto.Roles = roles

	return dto
}

func createRoles() {
	roleRepository := repository.NewRoleRepository(db.Conn.GetPgsqTemplate())
	roles := []sqlDomain.Role{
		{
			Name: sqlDomain.ADMIN,
		}, {
			Name: sqlDomain.USER,
		},
	}
	_, err := roleRepository.FindAll(0, 10)
	if err != nil {
		for _, i := range roles {
			roleRepository.Save(i)
		}
	}
}
