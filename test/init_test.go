package test

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var (
	MyToken     string
	testEnginer *gin.Engine
)

func getToken() string {

	token, err := util.SignUp(testEnginer, domain.USER, domain.ADMIN, domain.USER)
	if err != nil {
		panic(err)
	}

	return token
}

func init() {
	util.RemoveDatabase()
	testEnginer = gin.New()
	MyToken = getToken()
}
