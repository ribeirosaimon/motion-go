package middleware

import (
	sqlDomain2 "github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/exceptions"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func Authorization(roles ...sqlDomain2.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		motionValues := c.GetHeader("MotionRole")

		if authHeader != "" {
			authToken := strings.Split(authHeader, " ")
			if len(authToken) == 2 && authToken[0] == "Bearer" {
				bearerToken := authToken[1]
				// verify if exist Session for this user
				savedSession, err := repository.NewSessionRepository(db.Conn.GetPgsqTemplate()).
					FindByField("id", bearerToken)

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
					notAutenticateEmailSyncUser(c, profile)
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
				notAutenticateEmailSyncUser(c, profile)
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

func notAutenticateEmailSyncUser(c *gin.Context, profile sqlDomain2.Profile) {
	if profile.Status == sqlDomain2.ACTIVE || strings.Contains(c.Request.RequestURI, "validate") ||
		strings.Contains(c.Request.RequestURI, "whoami") {
		return
	}
	exceptions.Forbidden().Throw(c)
}

func putLoggedUserInContext(c *gin.Context, roleLoggedUser sqlDomain2.Role, p sqlDomain2.Profile, s sqlDomain2.Session) {
	var loggedUser LoggedUser
	loggedUser.ProfileId = p.Id
	loggedUser.Name = p.Name
	loggedUser.Role = roleLoggedUser
	loggedUser.SessionId = s.Id

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
	Name      string          `json:"name"`
	ProfileId uint64          `json:"loggedId"`
	SessionId string          `json:"sessionId"`
	Role      sqlDomain2.Role `json:"role"`
}
