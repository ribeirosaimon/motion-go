package test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/baseapp/pkg/service"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/internal/util"
)

func PutUserInContext(ctx *gin.Context, role sqlDomain.RoleEnum) sqlDomain.Session {
	propertiesFile := "config.test.properties"

	gin.SetMode(gin.TestMode)
	rootDir, _ := util.FindRootDir()
	dir := fmt.Sprintf("%s/%s", rootDir, propertiesFile)

	db.Conn = &db.Connections{}
	db.Conn.InitializeTestDatabases(properties.MustLoadFile(dir, properties.UTF8))

	setUpRoles()
	var user = sqlDomain.MotionUser{
		Name:     "test",
		Email:    "test@test.com",
		Password: "123",
	}
	profileUser, err := service.NewProfileService(db.Conn).SaveProfileUser(user, role)
	if err != nil {
		panic("error in profileUser")
	}
	session, err := service.NewSessionService(db.Conn).SaveUserSession(profileUser)
	if err != nil {
		panic("error in sessionUser")
	}
	var loggedUser = middleware.LoggedUser{
		Name:   profileUser.Name,
		UserId: profileUser.User.Id,
	}
	ctx.Set("loggedUser")
	return session
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
