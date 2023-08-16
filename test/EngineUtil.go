package test

import (
	"fmt"
	"github.com/ribeirosaimon/motion-go/scraping"

	"github.com/ribeirosaimon/motion-go/baseapp/pkg/dto"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/service"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/internal/util"
)

func SetUpTest(ctx *gin.Context, role sqlDomain.RoleEnum) middleware.LoggedUser {
	propertiesFile := "config.test.properties"

	// var err error
	gin.SetMode(gin.TestMode)

	rootDir, _ := util.FindRootDir()
	dir := fmt.Sprintf("%s/%s", rootDir, propertiesFile)

	db.Conn = &db.Connections{}
	db.Conn.InitializeTestDatabases(properties.MustLoadFile(dir, properties.UTF8))

	middleware.Cache = &middleware.MotionCache{
		Company: make(map[string]*middleware.Store),
		Service: scraping.NewScrapingService(db.Conn),
	}

	setUpRoles()
	var roles []sqlDomain.RoleEnum

	roles = append(roles, role)

	var loginDto = dto.LoginDto{
		Email:    "teste@teste.com",
		Password: "teste",
	}

	var signUp = dto.SignUpDto{
		Name:     "teste",
		Roles:    roles,
		LoginDto: loginDto,
	}

	userRepository := repository.NewUserRepository(db.Conn.GetPgsqTemplate())
	loginService := service.NewAuthService(db.Conn)
	var profile sqlDomain.Profile

	if !userRepository.ExistByField("email", loginDto.Email) {
		profile, _ = loginService.SignUp(signUp)
	} else {
		profileRepository := repository.NewProfileRepository(db.Conn.GetPgsqTemplate())
		user, _ := userRepository.FindByField("email", signUp.Email)
		profile, _ = profileRepository.FindByField("motion_user_id", user.Id)
	}

	sessionService := service.NewSessionService(db.Conn)

	sessionService.SaveUserSession(profile)

	var loggedUser = middleware.LoggedUser{
		Name:   profile.Name,
		UserId: profile.MotionUserId,
	}
	ctx.Set("loggedUser", loggedUser)
	return loggedUser
}

func setUpRoles() {
	connect := db.Conn.GetPgsqTemplate()
	defer db.Conn.ClosePostgreSQL()

	roleRepository := repository.NewRoleRepository(connect)
	allRoles := []sqlDomain.RoleEnum{sqlDomain.USER, sqlDomain.ADMIN}
	for _, i := range allRoles {
		existByName := roleRepository.ExistByField("name", i)
		if !existByName {
			roleRepository.Save(sqlDomain.Role{Name: i})
		}

	}
}

type SignUpTestDto struct {
	loginTestDto
	Name  string               `json:"name"`
	Roles []sqlDomain.RoleEnum `json:"roles"`
}

type loginTestDto struct {
	Email      string             `json:"email"`
	Password   string             `json:"password,omitempty"`
	Token      string             `json:"token"`
	LoggedId   uint32             `json:"id"`
	LoggedRole sqlDomain.RoleEnum `json:"loggedRole"`
}
