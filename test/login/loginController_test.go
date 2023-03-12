package login

import (
	"testing"

	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestLoginController(t *testing.T) {
	user := util.SetUpToTest()

	_, err := util.CreateLoginInApi(user.Name, user.Password)
	if err != nil {
		t.Error(err)
	}
}
