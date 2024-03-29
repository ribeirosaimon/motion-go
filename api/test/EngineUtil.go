package test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/api/internal/config"
	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"github.com/ribeirosaimon/motion-go/api/internal/dto"
	"github.com/ribeirosaimon/motion-go/api/internal/middleware"
	"github.com/ribeirosaimon/motion-go/api/internal/repository"
	"github.com/ribeirosaimon/motion-go/api/internal/service"
	"github.com/ribeirosaimon/motion-go/api/src/scraping"
	"github.com/ribeirosaimon/motion-go/confighub/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/confighub/util"
)

func SetUpTest(ctx *gin.Context, role sqlDomain.RoleEnum) middleware.LoggedUser {
	propertiesFile := "config.test.properties"

	// var err error
	gin.SetMode(gin.TestMode)

	rootDir, _ := util.FindRootDir()
	dir := fmt.Sprintf("%s/%s", rootDir, propertiesFile)

	db.Conn = &db.Connections{}
	propertiesTestFile := properties.MustLoadFile(dir, properties.UTF8)
	db.Conn.InitializeTestDatabases(propertiesTestFile)

	motionConfig := config.NewMotionConfig(propertiesTestFile)
	middleware.Cache = &middleware.MotionCache{
		Company: make(map[string]*middleware.Store),
		Service: scraping.NewScrapingService(db.Conn),
		Config:  motionConfig,
	}

	config.GetMotionConfig()
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

	session, err := sessionService.SaveUserSession(profile)
	if err != nil {
		panic(err)
	}

	var loggedUser = middleware.LoggedUser{
		Name:      profile.Name,
		ProfileId: profile.MotionUserId,
		SessionId: session.Id,
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
