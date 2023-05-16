package test

import (
	"fmt"
	"net/http"
	"path"
	"runtime"

	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

var (
	tokenUser   string
	tokenAdmin  string
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
func UpdateAdminToken(req *http.Request) {
	if tokenAdmin == "" {
		tokenAdmin, _ = SignUp(sqlDomain.USER, sqlDomain.ADMIN)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenAdmin))
	req.Header.Add("MotionRole", string(sqlDomain.ADMIN))
}

func UpdateUserToken(req *http.Request) {
	if tokenUser == "" {
		tokenUser, _ = SignUp(sqlDomain.USER, sqlDomain.ADMIN)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenUser))
	req.Header.Add("MotionRole", string(sqlDomain.USER))

}

func init() {
	CreateEnginer()
}
