package security

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/pkg/exceptions"
	"golang.org/x/crypto/bcrypt"
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

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, storedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	return err
}
