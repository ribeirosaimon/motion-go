package test

import (
	"fmt"
	"testing"

	"github.com/ribeirosaimon/motion-go/domain/sql"
	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestLoginAndSignUpController(t *testing.T) {
	session, err := util.SignUp(testEnginer, sql.USER, sql.ADMIN, sql.USER)

	if err != nil {
		panic(err)
	}
	util.SuccessTest(fmt.Sprintf("Expeted one session id: %s", session))
}
