package login

import (
	"fmt"
	"testing"

	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestLoginAndSignUpController(t *testing.T) {
	defer util.RemoveDatabase()
	session, err := util.SignUp()

	if err != nil {
		util.ErrorTest(err.Error())
	}
	util.SuccessTest(fmt.Sprintf("Expeted one session id: %s", session.SessionId))
}
