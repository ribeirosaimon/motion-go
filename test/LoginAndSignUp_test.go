package test

import (
	"fmt"
	"testing"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestLoginAndSignUpController(t *testing.T) {
	defer util.RemoveDatabase()
	session, err := util.SignUp(domain.USER, domain.ADMIN, domain.USER)

	if err != nil {
		util.ErrorTest(err.Error())
	}
	util.SuccessTest(fmt.Sprintf("Expeted one session id: %s", session))
}
