package login

import (
	"fmt"
	"testing"

	"github.com/ribeirosaimon/motion-go/test/util"
)

func TestLoginController(t *testing.T) {
	defer util.RemoveDatabase()
	sesion, err := util.SignUp()

	if err != nil {
		panic(err)
	}

	fmt.Println(sesion)

}
