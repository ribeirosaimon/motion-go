package test

//
// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"math/rand"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"strings"
// 	"time"
//
// 	"github.com/ribeirosaimon/motion-go/baseapp/pkg/router"
// 	"github.com/ribeirosaimon/motion-go/internal/db"
// 	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
// 	"github.com/ribeirosaimon/motion-go/internal/repository"
// )
//
// var (
// 	testJwtToken string
// 	user         SignUpTestDto
// )
//
// func Token(roles ...sqlDomain.RoleEnum) (string, SignUpTestDto) {
// 	if testJwtToken != "" {
// 		return testJwtToken, user
// 	}
// 	return signUp(roles...)
// }
//
// func signUp(roles ...sqlDomain.RoleEnum) (string, SignUpTestDto) {
// 	user := createUser(roles...)
// 	jsonData, _ := json.Marshal(user)
//
// 	e := CreateEngine(router.NewLoginRouter)
//
// 	req, _ := http.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(jsonData))
//
// 	resp := httptest.NewRecorder()
// 	e.ServeHTTP(resp, req)
//
// 	var signProfileResponse = sqlDomain.Profile{}
// 	err := json.Unmarshal(resp.Body.Bytes(), &signProfileResponse)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	dto := loginTestDto{Email: user.Email, Password: user.Password}
// 	jsonLoginDto, err := json.Marshal(dto)
//
// 	loginResp := httptest.NewRecorder()
// 	loginReq, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(jsonLoginDto))
// 	e.ServeHTTP(loginResp, loginReq)
// 	return strings.Replace(string(loginResp.Body.Bytes()), "\"", "", -1), user
// }
//
// func createUser(roles ...sqlDomain.RoleEnum) SignUpTestDto {
//
// 	rand.Seed(time.Now().UnixNano())
// 	nameRandom := strconv.Itoa(rand.Intn(1000000))
// 	password := strconv.Itoa(rand.Intn(1000000))
// 	emailRandom := fmt.Sprintf("%s@email.com", strconv.Itoa(rand.Intn(1000000)))
//
// 	var dto SignUpTestDto
// 	dto.Name = nameRandom
// 	dto.Password = password
// 	dto.Email = emailRandom
//
// 	allRoles := []sqlDomain.Role{
// 		{
// 			Name: sqlDomain.ADMIN,
// 		}, {
// 			Name: sqlDomain.USER,
// 		},
// 	}
//
// 	roleRepository := repository.NewRoleRepository(db.Conn.GetPgsqTemplate())
//
// 	for _, v := range allRoles {
// 		if !roleRepository.ExistByField("name", v.Name) {
// 			roleRepository.Save(v)
// 		}
// 	}
//
// 	for _, v := range roles {
// 		r, _ := roleRepository.FindByField("name", v)
// 		dto.Roles = append(dto.Roles, r.Name)
// 	}
// 	return dto
// }
//
// type SignUpTestDto struct {
// 	loginTestDto
// 	Name  string               `json:"name"`
// 	Roles []sqlDomain.RoleEnum `json:"roles"`
// }
//
// type loginTestDto struct {
// 	Email      string             `json:"email"`
// 	Password   string             `json:"password,omitempty"`
// 	Token      string             `json:"token"`
// 	LoggedId   uint32             `json:"id"`
// 	LoggedRole sqlDomain.RoleEnum `json:"loggedRole"`
// }
