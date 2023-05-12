package profile

import (
	"database/sql"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/domain"
	sqlDomain2 "github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	repository2 "github.com/ribeirosaimon/motion-go/internal/repository"

	"gorm.io/gorm"
)

type Service struct {
	profileRepository *repository2.MotionSQLRepository[sqlDomain2.Profile]
	roleRepository    *repository2.MotionSQLRepository[sqlDomain2.Role]
	closeDb           *sql.DB
}

func NewProfileService(conn *gorm.DB, close *sql.DB) Service {
	return Service{
		profileRepository: repository2.NewProfileRepository(conn),
		roleRepository:    repository2.NewRoleRepository(conn),
		closeDb:           close,
	}

}
func (l Service) SaveProfileUser(user sqlDomain2.MotionUser, roles []sqlDomain2.RoleEnum) (sqlDomain2.Profile, error) {
	var profile sqlDomain2.Profile

	profile.Name = user.Name

	for _, role := range roles {
		field, err := l.roleRepository.FindByField("name", role)
		if err != nil {
			return sqlDomain2.Profile{}, err
		}
		profile.Roles = []sqlDomain2.Role{field}
	}

	profile.Status = domain.ACTIVE
	profile.Birthday = user.Birthday
	profile.FamilyName = user.LastName
	profile.CreatedAt = time.Now()
	profile.User = user
	profile.UserId = user.GetId().(uint64)

	save, err := l.profileRepository.Save(profile)
	if err != nil {
		return sqlDomain2.Profile{}, err
	}
	return save, nil
}

func (l Service) FindProfileByUserId(id uint64) (sqlDomain2.Profile, error) {
	byId, err := l.profileRepository.FindWithPreloads("Roles", id)
	if err != nil {
		return sqlDomain2.Profile{}, err
	}
	return byId, nil
}
