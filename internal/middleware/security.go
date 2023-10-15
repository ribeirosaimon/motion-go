package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func haveRole(role sqlDomain.Role, roles []sqlDomain.Role) bool {
	for _, findRole := range roles {
		if findRole == role {
			return true
		}
	}
	return false
}

func Authorization(roles ...sqlDomain.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		motionValues := c.GetHeader("MotionRole")

		if authHeader != "" {
			authToken := strings.Split(authHeader, " ")
			if len(authToken) == 2 && authToken[0] == "Bearer" {
				bearerToken := authToken[1]
				// verify if exist Session for this user
				savedSession, err := repository.NewSessionRepository(db.Conn.GetPgsqTemplate()).
					FindByField("session_id", bearerToken)

				if err != nil {
					exceptions.FieldError("this user does not exist").Throw(c)
					return
				}

				if roles != nil {
					// verify if exist this Role
					motionLoggedRole, err := repository.NewRoleRepository(db.Conn.GetPgsqTemplate()).
						FindByField("name", motionValues)
					if err != nil {
						exceptions.FieldError("you no have motion roles").Throw(c)
						return
					}
					// get Profile by sessionId
					profile, err := repository.NewProfileRepository(db.Conn.GetPgsqTemplate()).
						FindWithPreloads("Roles", savedSession.ProfileId)
					if err != nil {
						exceptions.Forbidden().Throw(c)
						return
					}
					// verify if profile have loggedRole send by header
					if !profile.HaveRole(motionLoggedRole.Name) || err != nil {
						exceptions.Forbidden().Throw(c)
						return
					}
					for _, v := range roles {
						if profile.HaveRole(v.Name) {
							putLoggedUserInContext(c, motionLoggedRole, profile, savedSession)
							c.Next()
							return
						}
					}
				}
				profile, err := repository.NewProfileRepository(db.Conn.GetPgsqTemplate()).
					FindWithPreloads("Roles", savedSession.ProfileId)
				putLoggedUserInContext(c, profile.Roles[0], profile, savedSession)
				c.Next()
				return

			} else {
				exceptions.Forbidden().Throw(c)
				c.Next()
			}
		}
		exceptions.Forbidden().Throw(c)
		c.Next()
	}
}

func putLoggedUserInContext(c *gin.Context, roleLoggedUser sqlDomain.Role, p sqlDomain.Profile, s sqlDomain.Session) {
	var loggedUser LoggedUser
	loggedUser.UserId = p.MotionUserId
	loggedUser.Name = p.Name
	loggedUser.Role = roleLoggedUser
	loggedUser.SessionId = s.SessionId

	c.Set("loggedUser", loggedUser)
}
func GetLoggedUser(c *gin.Context) (LoggedUser, error) {
	if _, ok := c.Get("loggedUser"); !ok {
		return LoggedUser{}, errors.New("not exist loggedUser")
	}
	return c.MustGet("loggedUser").(LoggedUser), nil
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
	Name      string         `json:"name"`
	UserId    uint64         `json:"loggedId"`
	SessionId string         `json:"sessionId"`
	Role      sqlDomain.Role `json:"role"`
}
