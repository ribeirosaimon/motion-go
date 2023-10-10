package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/src/pkg/dto"
	"github.com/ribeirosaimon/motion-go/src/pkg/service"
	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

func TestLoginController_SignUp(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	var signupDto = dto.SignUpDto{
		Name:  "testUser",
		Roles: []sqlDomain.RoleEnum{sqlDomain.USER},
		LoginDto: dto.LoginDto{
			Email:    "user@test.com",
			Password: "testPassword",
		},
	}

	configTest()
	jsonBytes, err := json.Marshal(signupDto)
	reader := bytes.NewReader(jsonBytes)

	c.Request = &http.Request{Body: ioutil.NopCloser(reader)}

	NewAuthController().SignUp(c)

	var response sqlDomain.Profile
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	roleRepository := repository.NewRoleRepository(db.Conn.GetPgsqTemplate())
	role, err := roleRepository.FindByField("name", sqlDomain.USER)
	assert.Equal(t, signupDto.Name, response.Name)
	assert.Equal(t, sqlDomain.USER, role.Name)
	assert.Equal(t, domain.ACTIVE, response.Status)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLoginController_LoginIsNotOk(t *testing.T) {

	w, c := configAuthTest()
	var loginDto = dto.LoginDto{
		Email:    "user@test.com",
		Password: "testPasswordIncorrect",
	}

	configTest()
	jsonBytes, err := json.Marshal(loginDto)
	reader := bytes.NewReader(jsonBytes)

	c.Request = &http.Request{Body: ioutil.NopCloser(reader)}
	NewAuthController().Login(c)

	var response exceptions.Error
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "bad field error: password", response.Message)
}

func TestLoginController_Login(t *testing.T) {
	var loginDto = dto.LoginDto{
		Email:    "user@test.com",
		Password: "testPassword",
	}

	userRepository := repository.NewUserRepository(db.Conn.GetPgsqTemplate())
	userBeforeLogin, err := userRepository.FindByField("email", loginDto.Email)

	if err != nil {
		TestLoginController_SignUp(t)
	}

	w, c := configAuthTest()

	configTest()
	jsonBytes, err := json.Marshal(loginDto)
	reader := bytes.NewReader(jsonBytes)

	c.Request = &http.Request{Body: ioutil.NopCloser(reader)}
	NewAuthController().Login(c)

	var response string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	sessionRepository := repository.NewSessionRepository(db.Conn.GetPgsqTemplate())
	session, err := sessionRepository.FindByField("session_id", response)
	profile, err := service.NewProfileService(db.Conn).FindProfileByUserId(session.ProfileId)

	userAfterLogin, err := userRepository.FindByField("email", loginDto.Email)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, session.SessionId, response)
	assert.Equal(t, profile.Id, session.ProfileId)
	assert.Equal(t, userBeforeLogin.LoginCount, userAfterLogin.LoginCount-1)
}

func TestLoginController_WhoAmI(t *testing.T) {

}

func configAuthTest() (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	test.SetUpTest(c, sqlDomain.USER)

	return w, c
}
