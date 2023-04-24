package test

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var (
	TokenUser   string
	TokenAdmin  string
	testEnginer *gin.Engine
)

func getToken(role sqlDomain.RoleEnum) string {
	var token string
	var err error

	if role == sqlDomain.USER {
		token, err = util.SignUp(testEnginer, sqlDomain.USER, sqlDomain.USER)
	} else if role == sqlDomain.ADMIN {
		token, err = util.SignUp(testEnginer, sqlDomain.ADMIN, sqlDomain.ADMIN)
	}

	if err != nil {
		panic(err)
	}

	return token
}

func init() {
	util.RemoveDatabase()
	testEnginer = gin.New()
	TokenUser = getToken(sqlDomain.USER)
	TokenAdmin = getToken(sqlDomain.ADMIN)
}
