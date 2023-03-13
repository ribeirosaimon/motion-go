package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ribeirosaimon/motion-go/repository"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/login"
)

var enginer *gin.Engine

func GetEnginer() *gin.Engine {
	return enginer
}
func CreateEngineRequest(method, path string, body io.Reader,
	controller func(*gin.Engine), session string) (
	*httptest.ResponseRecorder, *http.Request, error) {
	ginEngine := gin.New()
	controller(ginEngine)

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, nil, err
	}
	if session != "" {
		req.Header.Add("Authorization", session)
		req.Header.Add("MotionRole", "ADMIN")
	}

	w := httptest.NewRecorder()
	ginEngine.ServeHTTP(w, req)

	return w, req, nil
}

func SignUp() (domain.Session, error) {
	user := CreateUser()
	//jsonValues := make(map[string]string)
	//
	//jsonValues["email"] = user.Email
	//jsonValues["name"] = user.Name
	//jsonValues["password"] = user.Password

	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	sigUpResponse, _, err := CreateEngineRequest(http.MethodPost, "/sign-up",
		bytes.NewReader(jsonData), login.NewLoginController, "")
	var signProfileResponse = domain.Profile{}
	err = json.Unmarshal(sigUpResponse.Body.Bytes(), &signProfileResponse)
	if err != nil {
		panic(err)
	}
	dto := login.LoginDto{Email: user.Email, Password: user.Password}
	jsonLoginDto, err := json.Marshal(dto)
	resp, _, err := CreateEngineRequest(http.MethodPost, "/login",
		bytes.NewReader(jsonLoginDto), login.NewLoginController, "")

	var response = domain.Session{}
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		return domain.Session{}, err
	}

	return response, nil
}

func CreateUser() login.SignUpDto {
	createRoles()
	rand.Seed(time.Now().UnixNano())
	nameRandom := strconv.Itoa(rand.Intn(1000000))
	password := strconv.Itoa(rand.Intn(1000000))
	emailRandom := fmt.Sprintf("%s@email.com", strconv.Itoa(rand.Intn(1000000)))
	//lastNameRandom := strconv.Itoa(rand.Intn(1000000))
	//LoginAttempRandom := uint8(rand.Intn(10))
	//loginCountRandom := uint64(rand.Intn(101))

	var dto login.SignUpDto
	dto.Name = nameRandom
	dto.Password = password
	dto.Email = emailRandom
	return dto
}

func createRoles() {
	roleRepository := repository.NewRoleRepository(ConnectDatabaseTest())
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
