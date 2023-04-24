package util

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	sql2 "github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/login"
	"github.com/ribeirosaimon/motion-go/repository"
)

func CreateEngineRequest(enginer *gin.Engine, method, path string, body io.Reader, session string,
	role sql2.RoleEnum) (
	*httptest.ResponseRecorder, *http.Request, error) {

	AddController(enginer, "/api/v1/auth", login.NewLoginRouter)

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, nil, err
	}
	if session != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", session))
		req.Header.Add("MotionRole", string(role))
	}

	w := httptest.NewRecorder()
	enginer.ServeHTTP(w, req)

	return w, req, nil
}

func SignUp(enginer *gin.Engine, loggedRole sql2.RoleEnum, roles ...sql2.RoleEnum) (string, error) {
	user := CreateUser(roles...)
	jsonData, err := json.Marshal(user)

	AddController(enginer, "/api/v1/auth", login.NewLoginRouter)

	if err != nil {
		panic(err)
	}
	sigUpResponse, _, err := CreateEngineRequest(enginer, http.MethodPost, "/api/v1/auth/sign-up",
		bytes.NewReader(jsonData), "", loggedRole)
	var signProfileResponse = sql2.Profile{}
	err = json.Unmarshal(sigUpResponse.Body.Bytes(), &signProfileResponse)
	if err != nil {
		panic(err)
	}
	dto := login.LoginDto{Email: user.Email, Password: user.Password}
	jsonLoginDto, err := json.Marshal(dto)
	resp, _, err := CreateEngineRequest(enginer, http.MethodPost, "/api/v1/auth/login",
		bytes.NewReader(jsonLoginDto), "", loggedRole)

	return strings.Replace(string(resp.Body.Bytes()), "\"", "", -1), nil
}

func CreateUser(roles ...sql2.RoleEnum) login.SignUpDto {
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
	test, close := ConnectDatabaseTest()
	defer close.Close()
	roleRepository := repository.NewRoleRepository(test)
	roles := []sql2.Role{
		{
			Name: sql2.ADMIN,
		}, {
			Name: sql2.USER,
		},
	}
	_, err := roleRepository.FindAll(0, 10)
	if err != nil {
		for _, i := range roles {
			roleRepository.Save(i)
		}
	}
}

func SuccessTest(info string) string {
	return fmt.Sprintf("\033[32mSuccess:\033[0m %s.\"", info)
}

//
// func AssertEquals(t *testing.T, v1, v2 interface{}) {
//	// Obtém informações sobre a chamada anterior
//	_, _, line, _ := runtime.Caller(1)
//
//	assert.Equal(t, v1, v2)
//	if v1 != v2 {
//		t.Errorf("\033[31mError in line %d:\033[0m.\" Expected: %s but received: %s\n", line, v1, v2)
//	}
// }
//
// func AssertNotEquals(t *testing.T, v1, v2 interface{}) {
//	// Obtém informações sobre a chamada anterior
//	_, _, line, _ := runtime.Caller(1)
//
//	if v1 == v2 {
//		t.Errorf("\033[31mError in line %d:\033[0m.\" Expected: %s but received: %s\n", line, v1, v2)
//	}
// }

func AddController(enginer *gin.Engine, subs string, f func(engine *gin.RouterGroup,
	conn func() (*gorm.DB, *sql.DB))) {
	contains := false
	for _, route := range enginer.Routes() {
		if strings.Contains(route.Path, subs) {
			contains = true
			break
		}
	}
	if !contains {
		f(enginer.Group("/api/v1"), ConnectDatabaseTest)
	}
}
