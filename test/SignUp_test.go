package test

import (
	"fmt"
	"testing"

	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestLoginAndSignUpController(t *testing.T) {
	session, err := util.SignUp(testEnginer, sqlDomain.USER, sqlDomain.ADMIN, sqlDomain.USER)

	if err != nil {
		panic(err)
	}
	util.SuccessTest(fmt.Sprintf("Expeted one session id: %s", session))
}
