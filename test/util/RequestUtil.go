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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/health"
	"github.com/ribeirosaimon/motion-go/pkg/login"
	"github.com/ribeirosaimon/motion-go/repository"
)

var enginer *gin.Engine

func CreateEngineRequest(method, path string, body io.Reader) (
	*httptest.ResponseRecorder, *http.Request, error) {
	enginer = gin.New()

	health.NewHeathController(enginer)

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, nil, err
	}
	w := httptest.NewRecorder()
	enginer.ServeHTTP(w, req)

	return w, req, nil
}

func GetEnginer() *gin.Engine {
	return enginer
}

func SetUpToTest() domain.MotionUser {
	user := CreateUser()
	var userRepository = repository.NewUserRepository(ConnectDatabaseTest())
	save, err := userRepository.Save(user)
	if err != nil {
		panic(err)
	}
	return save
}

func CreateLoginInApi(email, password string) (domain.Session, error) {
	dto := login.LoginDto{Email: email, Password: password}
	jsonData, err := json.Marshal(dto)
	if err != nil {
		panic(err)
	}
	resp, _, err := CreateEngineRequest(http.MethodPost, "/login", bytes.NewReader(jsonData))

	var response = domain.Session{}
	err = json.Unmarshal(resp.Body.Bytes(), &response)

	if err != nil {
		return domain.Session{}, err
	}

	return response, nil
}

func CreateUser() domain.MotionUser {
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
