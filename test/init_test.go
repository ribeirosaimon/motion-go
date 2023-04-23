package test

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain/sql"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var (
	TokenUser   string
	TokenAdmin  string
	testEnginer *gin.Engine
)

func getToken(role sql.RoleEnum) string {
	var token string
	var err error

	if role == sql.USER {
		token, err = util.SignUp(testEnginer, sql.USER, sql.USER)
	} else if role == sql.ADMIN {
		token, err = util.SignUp(testEnginer, sql.ADMIN, sql.ADMIN)
	}

	if err != nil {
		panic(err)
	}

	return token
}

func init() {
	util.RemoveDatabase()
	testEnginer = gin.New()
	TokenUser = getToken(sql.USER)
	TokenAdmin = getToken(sql.ADMIN)
}
