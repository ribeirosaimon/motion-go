package security

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/database"
	"github.com/ribeirosaimon/motion-go/pkg/exceptions"
	"github.com/ribeirosaimon/motion-go/pkg/profile"
	"github.com/ribeirosaimon/motion-go/repository"
	"golang.org/x/crypto/bcrypt"
)

func haveRole(role domain.Role, roles []domain.Role) bool {
	for _, findRole := range roles {
		if findRole == role {
			return true
		}
	}
	return false
}

func Authorization(roles ...domain.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		connect, close := database.Connect()
		defer close.Close()
		authHeader := c.GetHeader("Authorization")
		motionValues := c.GetHeader("MotionRole")

		if authHeader != "" {
			authToken := strings.Split(authHeader, " ")
			if len(authToken) == 2 && authToken[0] == "Bearer" {
				bearerToken := authToken[1]
				// verify if exist Session for this user
				savedSession, err := repository.NewSessionRepository(connect).FindByField("session_id", bearerToken)
				// verify if exist this Role
				motionLoggedRole, err := repository.NewRoleRepository(connect).FindByField("name", motionValues)
				// get Profile by sessionId
				profile, err := profile.NewProfileService().FindProfileByUserId(savedSession.ProfileId)
				if err != nil {
					exceptions.Unauthorized(c)
					return
				}
				// verify if profile have loggedRole send by heder
				if !profile.HaveRole(motionLoggedRole.Name) || err != nil {
					exceptions.Unauthorized(c)
					return
				}
				for _, v := range roles {
					if profile.HaveRole(v.Name) {
						c.Next()
						return
					}
				}
			} else {
				exceptions.Unauthorized(c)
			}
		}
		exceptions.Unauthorized(c)
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
