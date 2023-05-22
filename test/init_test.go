package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

var (
	tokenUser   string
	tokenAdmin  string
	testEnginer *config.MotionGo
	ctx         *gin.Context
	recorder    *httptest.ResponseRecorder
)

func CreateEnginer() {
	propertiesFile := "config.test.properties"

	// gin.SetMode(gin.TestMode)
	testEnginer = config.NewMotionGo(propertiesFile, true)

	recorder = httptest.NewRecorder()
	ctx, testEnginer.MotionEngine = gin.CreateTestContext(recorder)

	db.Conn = &db.Connections{}
	db.Conn.InitializeTestDatabases(getCurrentDirectory())

	setUpRoles()

}

func AddRouter(v ...config.RoutersVersion) {
	testEnginer.AddRouter(v...)
	testEnginer.CreateRouters()
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
func AddAdminTokenInReq(req *http.Request) {
	if tokenAdmin == "" {
		tokenAdmin, _ = SignUp(sqlDomain.USER, sqlDomain.ADMIN)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenAdmin))
	req.Header.Add("MotionRole", string(sqlDomain.ADMIN))
}

func AddUserTokenInReq(req *http.Request) {
	if tokenUser == "" {
		tokenUser, _ = SignUp(sqlDomain.USER, sqlDomain.ADMIN)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenUser))
	req.Header.Add("MotionRole", string(sqlDomain.USER))

}

func Context() *gin.Context {
	return ctx
}

func HttpResponse() *httptest.ResponseRecorder {
	returnedValue := recorder
	recorder = httptest.NewRecorder()
	return returnedValue
}
func init() {
	CreateEnginer()
}
