package security

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/exceptions"
	"github.com/ribeirosaimon/motion-go/pkg/profile"
	"github.com/ribeirosaimon/motion-go/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func haveRole(role domain.Role, roles []domain.Role) bool {
	for _, findRole := range roles {
		if findRole == role {
			return true
		}
	}
	return false
}

func Authorization(dbConn func() (*gorm.DB, *sql.DB), roles ...domain.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		connect, close := dbConn()
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
				if err != nil {
					exceptions.FieldError("you no have motion roles").Throw(c)
					return
				}
				// get Profile by sessionId
				profile, err := profile.NewProfileService(connect, close).FindProfileByUserId(savedSession.ProfileId)
				if err != nil {
					exceptions.Forbidden().Throw(c)
					return
				}
				// verify if profile have loggedRole send by heder
				if !profile.HaveRole(motionLoggedRole.Name) || err != nil {
					exceptions.Forbidden().Throw(c)
					return
				}
				for _, v := range roles {
					if profile.HaveRole(v.Name) {
						putLoggedUserInContext(c, motionLoggedRole, profile)
						c.Next()
						return
					}
				}
			} else {
				exceptions.Forbidden().Throw(c)
			}
		}
		exceptions.Forbidden().Throw(c)
	}
}

func putLoggedUserInContext(c *gin.Context, roleLoggedser domain.Role, p domain.Profile) {
	var loggedUser LoggedUser
	loggedUser.UserId = p.UserId
	loggedUser.Name = p.Name
	loggedUser.Role = roleLoggedser

	c.Set("loggedUser", loggedUser)
}
func GetLoggedUser(c *gin.Context) LoggedUser {
	return c.MustGet("loggedUser").(LoggedUser)
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

type LoggedUser struct {
	Name   string      `json:"name"`
	UserId uint64      `json:"loggedId"`
	Role   domain.Role `json:"role"`
}
