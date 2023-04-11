package test

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var (
	TokenUser   string
	TokenAdmin  string
	testEnginer *gin.Engine
)

func getToken(role domain.RoleEnum) string {
	var token string
	var err error

	if role == domain.USER {
		token, err = util.SignUp(testEnginer, domain.USER, domain.USER)
	} else if role == domain.ADMIN {
		token, err = util.SignUp(testEnginer, domain.ADMIN, domain.ADMIN)
	}

	if err != nil {
		panic(err)
	}

	return token
}

func init() {
	util.RemoveDatabase()
	testEnginer = gin.New()
	TokenUser = getToken(domain.USER)
	TokenAdmin = getToken(domain.ADMIN)
}
