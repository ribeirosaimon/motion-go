package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlDomain2 "github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/dto"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/internal/service"
	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

func TestLoginController_SignUp(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	var signupDto = dto.SignUpDto{
		Name:  "testUser",
		Roles: []sqlDomain2.RoleEnum{sqlDomain2.USER},
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

	var response sqlDomain2.Profile
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}

	roleRepository := repository.NewRoleRepository(db.Conn.GetPgsqTemplate())

	role, err := roleRepository.FindByField("name", sqlDomain2.USER)
	assert.Equal(t, signupDto.Name, response.Name)
	assert.Equal(t, response.Status, sqlDomain2.EMAIL_SYNC)

	assert.Equal(t, sqlDomain2.USER, role.Name)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLoginController_ValidateEmail(t *testing.T) {

	configTest()

	connection := db.Conn.GetPgsqTemplate()

	userRepository := repository.NewUserRepository(connection)
	user, err := userRepository.FindByField("email", "teste@teste.com")

	profileRepository := repository.NewProfileRepository(connection)
	profile, err := profileRepository.FindByField("motion_user_id", user.Id)

	if err != nil {
		panic(err)
	}

	newRecorder := httptest.NewRecorder()
	newContext, _ := gin.CreateTestContext(newRecorder)
	loggedUser := test.SetUpTest(newContext, sqlDomain2.USER)

	jsonBytes, err := json.Marshal(dto.ValidateEmailDto{Code: profile.Code})
	reader := bytes.NewReader(jsonBytes)

	newContext.Request = &http.Request{Body: ioutil.NopCloser(reader)}

	NewAuthController().ValidateEmail(newContext)
	transactionRepository := repository.NewTransactionRepository(connection)

	transaction, err := transactionRepository.FindByField("session_id", loggedUser.SessionId)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, newRecorder.Code)
	assert.Equal(t, transaction.SessionId, loggedUser.SessionId)
	assert.Equal(t, transaction.ProfileId, loggedUser.ProfileId)
	assert.Equal(t, transaction.OperationType, sqlDomain2.DEPOSIT)
}

func TestLoginController_ValidateEmailOnlyOneTime(t *testing.T) {

	configTest()
	connection := db.Conn.GetPgsqTemplate()

	userRepository := repository.NewUserRepository(connection)
	user, err := userRepository.FindByField("email", "teste@teste.com")

	profileRepository := repository.NewProfileRepository(connection)
	profile, err := profileRepository.FindByField("motion_user_id", user.Id)

	if err != nil {
		panic(err)
	}

	newRecorder := httptest.NewRecorder()
	newContext, _ := gin.CreateTestContext(newRecorder)

	jsonBytes, err := json.Marshal(dto.ValidateEmailDto{Code: profile.Code})
	reader := bytes.NewReader(jsonBytes)

	newContext.Request = &http.Request{Body: ioutil.NopCloser(reader)}

	NewAuthController().ValidateEmail(newContext)

	newRecorder = httptest.NewRecorder()
	newContext, _ = gin.CreateTestContext(newRecorder)
	NewAuthController().ValidateEmail(newContext)

	assert.Equal(t, http.StatusBadRequest, newRecorder.Code)
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
	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, "Not Found", response.Message)
}

func TestLoginController_Login(t *testing.T) {

	token, w, _ := loginAndGetToken(t)

	sessionRepository := repository.NewSessionRepository(db.Conn.GetPgsqTemplate())
	session, err := sessionRepository.FindByField("id", token)
	if err != nil {
		panic(err)
	}
	userRepository := repository.NewUserRepository(db.Conn.GetPgsqTemplate())
	profile, err := service.NewProfileService(db.Conn).FindProfileByUserId(session.ProfileId)

	userAfterLogin, err := userRepository.FindByField("email", "user@test.com")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, session.Id, token)
	assert.Equal(t, profile.Id, session.ProfileId)
	assert.NotEqual(t, 0, userAfterLogin.LoginCount)
}

func configAuthTest() (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	test.SetUpTest(c, sqlDomain2.USER)

	return w, c
}

func loginAndGetToken(t *testing.T) (string, *httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var loginDto = dto.LoginDto{
		Email:    "user@test.com",
		Password: "testPassword",
	}
	configTest()
	userRepository := repository.NewUserRepository(db.Conn.GetPgsqTemplate())
	_, err := userRepository.FindByField("email", loginDto.Email)

	if err != nil {
		TestLoginController_SignUp(t)
	}

	jsonBytes, err := json.Marshal(loginDto)
	reader := bytes.NewReader(jsonBytes)

	c.Request = &http.Request{Body: ioutil.NopCloser(reader)}
	NewAuthController().Login(c)

	var response string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error in unmarshal json %d", w.Body)
	}
	return response, w, c
}
