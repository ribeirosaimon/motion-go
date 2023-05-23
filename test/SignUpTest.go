package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/login"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

var testJwtToken string

func Token() string {
	if testJwtToken != "" {
		return testJwtToken
	}
	return signUp()
}

func signUp() string {
	user := createUser()
	jsonData, _ := json.Marshal(user)

	e := CreateEngine(login.NewLoginRouter)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/sign-up", bytes.NewReader(jsonData))

	resp := httptest.NewRecorder()
	e.ServeHTTP(resp, req)

	var signProfileResponse = sqlDomain.Profile{}
	err := json.Unmarshal(resp.Body.Bytes(), &signProfileResponse)
	if err != nil {
		panic(err)
	}

	dto := loginDto{Email: user.Email, Password: user.Password}
	jsonLoginDto, err := json.Marshal(dto)

	loginResp := httptest.NewRecorder()
	loginReq, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(jsonLoginDto))
	e.ServeHTTP(loginResp, loginReq)
	return strings.Replace(string(loginResp.Body.Bytes()), "\"", "", -1)
}

func createUser() signUpDto {
	createRoles()
	rand.Seed(time.Now().UnixNano())
	nameRandom := strconv.Itoa(rand.Intn(1000000))
	password := strconv.Itoa(rand.Intn(1000000))
	emailRandom := fmt.Sprintf("%s@email.com", strconv.Itoa(rand.Intn(1000000)))

	var dto signUpDto
	dto.Name = nameRandom
	dto.Password = password
	dto.Email = emailRandom
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

type signUpDto struct {
	loginDto
	Name  string               `json:"name"`
	Roles []sqlDomain.RoleEnum `json:"roles"`
}

type loginDto struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
