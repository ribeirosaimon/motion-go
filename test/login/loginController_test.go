package login

import (
	"testing"

	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestLoginAndSignUpController(t *testing.T) {
	defer util.RemoveDatabase()
	_, err := util.SignUp()

	if err != nil {
		t.Errorf(err.Error())
	}

}
