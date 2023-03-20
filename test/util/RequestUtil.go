package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/login"
	"github.com/ribeirosaimon/motion-go/repository"
)

func CreateEngineRequest(method, path string, body io.Reader,
	controller func(group *gin.RouterGroup), session string, role domain.RoleEnum) (
	*httptest.ResponseRecorder, *http.Request, error) {
	enginer := gin.New()
	controller(enginer.Group(""))

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

func SignUp(loggedRole domain.RoleEnum, roles ...domain.RoleEnum) (string, error) {
	user := CreateUser(roles...)
	jsonData, err := json.Marshal(user)
	var loginRouter = func(engine *gin.RouterGroup) {
		login.NewLoginRouter(engine.Group("/api/v1"), ConnectDatabaseTest)
	}
	if err != nil {
		panic(err)
	}
	sigUpResponse, _, err := CreateEngineRequest(http.MethodPost, "/api/v1/auth/sign-up",
		bytes.NewReader(jsonData), loginRouter, "", loggedRole)
	var signProfileResponse = domain.Profile{}
	err = json.Unmarshal(sigUpResponse.Body.Bytes(), &signProfileResponse)
	if err != nil {
		panic(err)
	}
	dto := login.LoginDto{Email: user.Email, Password: user.Password}
	jsonLoginDto, err := json.Marshal(dto)
	resp, _, err := CreateEngineRequest(http.MethodPost, "/api/v1/auth/login",
		bytes.NewReader(jsonLoginDto), loginRouter, "", loggedRole)

	return strings.Replace(string(resp.Body.Bytes()), "\"", "", -1), nil
}

func CreateUser(roles ...domain.RoleEnum) login.SignUpDto {
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
	roles := []domain.Role{
		{
			Name: domain.ADMIN,
		}, {
			Name: domain.USER,
		},
	}
	_, err := roleRepository.FindAll(0, 10)
	if err != nil {
		for _, i := range roles {
			roleRepository.Save(i)
		}
	}
}

func SuccessTest(info string) {
	fmt.Println(fmt.Sprintf("\033[32mSuccess:\033[0m %s.\"", info))
}

func ErrorTest(info string) {
	fmt.Println(fmt.Sprintf("\033[31mError:\033[0m %s.\"", info))
}
