package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/config"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/internal/util"
)

func PerformRequest(r http.Handler, method, path, role string, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)

	req.Header.Set("Content-Type", "application/json")
	if role != "" {

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token()))
		if role == "" {
			req.Header.Set("MotionRole", "ADMIN")
		} else {
			req.Header.Set("MotionRole", role)
		}

	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func CreateEngine(controller func() config.MotionController) *gin.Engine {
	propertiesFile := "config.test.properties"

	gin.SetMode(gin.TestMode)
	rootDir, _ := util.FindRootDir()
	dir := fmt.Sprintf("%s/%s", rootDir, propertiesFile)

	db.Conn = &db.Connections{}
	db.Conn.InitializeTestDatabases(rootDir)

	setUpRoles()

	motion := config.NewMotionGo(dir)

	c := controller()
	group := motion.MotionEngine.Group(c.Path)
	for _, v := range c.Handlers {
		v.Middleware = append(v.Middleware, v.Service)

		group.Handle(v.Method, v.Path, v.Middleware...)
	}
	return motion.MotionEngine
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
