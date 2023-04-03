package test

import (
	"fmt"
	"testing"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestLoginAndSignUpController(t *testing.T) {
	session, err := util.SignUp(testEnginer, domain.USER, domain.ADMIN, domain.USER)

	if err != nil {
		util.ErrorTest(err.Error())
	}
	util.SuccessTest(fmt.Sprintf("Expeted one session id: %s", session))
}
