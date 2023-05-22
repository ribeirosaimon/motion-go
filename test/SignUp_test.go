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
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

var loginVersion = config.RoutersVersion{
	Version: "v1",
	Handlers: []func() config.MotionController{
		login.NewLoginRouter,
	},
}

func SignUp(roles ...sqlDomain.RoleEnum) (string, error) {
	user := CreateUser(roles...)
	jsonData, _ := json.Marshal(user)

	AddRouter(loginVersion)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/sign-up", bytes.NewReader(jsonData))

	resp := httptest.NewRecorder()
	testEnginer.MotionEngine.ServeHTTP(resp, req)

	var signProfileResponse = sqlDomain.Profile{}
	err := json.Unmarshal(resp.Body.Bytes(), &signProfileResponse)
	if err != nil {
		panic(err)
	}

	dto := login.LoginDto{Email: user.Email, Password: user.Password}
	jsonLoginDto, err := json.Marshal(dto)

	loginResp := httptest.NewRecorder()
	loginReq, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(jsonLoginDto))
	testEnginer.MotionEngine.ServeHTTP(loginResp, loginReq)
	return strings.Replace(string(loginResp.Body.Bytes()), "\"", "", -1), nil
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
