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
		util.ErrorTest(err.Error())
	}

	return token
}

func init() {
	testEnginer = gin.New()
	util.RemoveDatabase()
	MyToken = getToken()

}
