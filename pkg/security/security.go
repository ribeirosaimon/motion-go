package security

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/exceptions"
)

type RoleEnum string

const (
	ADMIN RoleEnum = "ADMIN"
	USER  RoleEnum = "USER"
)

func haveRole(role RoleEnum, roles []RoleEnum) bool {
	for _, findRole := range roles {
		if findRole == role {
			return true
		}
	}
	return false
}

func Authorization(roles ...RoleEnum) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Next()

		if haveRole(USER, roles) {

			fmt.Println("you have permision")
		}
		exceptions.Unauthorized(c)
		fmt.Println("you not have permision")
	}
}
