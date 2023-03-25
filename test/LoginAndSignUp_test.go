package test

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var signUpEnginer = gin.New()

func TestLoginAndSignUpController(t *testing.T) {
	session, err := util.SignUp(signUpEnginer, domain.USER, domain.ADMIN, domain.USER)

	if err != nil {
		util.ErrorTest(err.Error())
	}
	util.SuccessTest(fmt.Sprintf("Expeted one session id: %s", session))
}
