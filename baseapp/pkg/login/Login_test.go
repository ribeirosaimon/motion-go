package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

var (
	userEmail    string
	userPassword string

	e = test.CreateEngine(NewLoginRouter)
)

func TestSignUp(t *testing.T) {

	user := createUser()
	jsonData, _ := json.Marshal(user)

	w := test.PerformRequest(e, http.MethodPost, "/auth/sign-up", bytes.NewReader(jsonData), nil)
	s := sqlDomain.Profile{}
	json.Unmarshal(w.Body.Bytes(), &s)

	assert.NotEqual(t, 0, s.Id, "Not to be 0")
	assert.NotEqual(t, "", s.Name, "Not to be clear")
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLogin(t *testing.T) {
	TestSignUp(t)
	dto := loginDto{Email: userEmail, Password: userPassword}
	jsonLoginDto, _ := json.Marshal(dto)
	w := test.PerformRequest(e, http.MethodPost, "/auth/login", bytes.NewReader(jsonLoginDto), nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func createUser() signUpDto {
	createRoles()
	rand.Seed(time.Now().UnixNano())
	nameRandom := strconv.Itoa(rand.Intn(1000000))
	password := strconv.Itoa(rand.Intn(1000000))
	emailRandom := fmt.Sprintf("%s@email.com", strconv.Itoa(rand.Intn(1000000)))

	var dto signUpDto
	dto.Name = nameRandom
	dto.Password, userPassword = password, password
	dto.Email, userEmail = emailRandom, emailRandom
	dto.Roles = []sqlDomain.RoleEnum{sqlDomain.ADMIN, sqlDomain.USER}

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
