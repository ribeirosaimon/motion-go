package test

import (
	"path"
	"runtime"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var (
	TokenUser   string
	TokenAdmin  string
	TestEnginer *config.MotionGo
)

func CreateEnginer() {
	propertiesFile := "config.test.properties"

	TestEnginer = config.NewMotionGo(propertiesFile)

	db.Conn = &db.Connections{}
	db.Conn.InitializeTestDatabases(getCurrentDirectory())

	setUpRoles()

}

func AddRouter(v ...config.RoutersVersion) {
	TestEnginer.AddRouter(v...)
	TestEnginer.CreateRouters()
}

func getCurrentDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get current file info")
	}
	dir := path.Dir(filename)
	return dir
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

func getToken(role sqlDomain.RoleEnum) string {
	var token string
	var err error

	if role == sqlDomain.USER {
		token, err = util.SignUp(sqlDomain.USER, sqlDomain.USER)
	} else if role == sqlDomain.ADMIN {
		// token, err = util.SignUp(testEnginer, sqlDomain.ADMIN, sqlDomain.ADMIN)
	}

	if err != nil {
		panic(err)
	}

	return token
}

func init() {
	CreateEnginer()
	// TestEnginer.CreateRouters(TestEnginer.PropertiesFile.GetInt("server.port.baseapp", 0))

	// TokenUser = getToken(sqlDomain.USER)
	// TokenAdmin = getToken(sqlDomain.ADMIN)
}
