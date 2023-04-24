package profile

import (
	"database/sql"
	"time"

	sql2 "github.com/ribeirosaimon/motion-go/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/repository"
	"gorm.io/gorm"
)

type Service struct {
	profileRepository repository.MotionRepository[sql2.Profile]
	roleRepository    repository.MotionRepository[sql2.Role]
	closeDb           *sql.DB
}

func NewProfileService(conn *gorm.DB, close *sql.DB) Service {
	return Service{
		profileRepository: repository.NewProfileRepository(conn),
		roleRepository:    repository.NewRoleRepository(conn),
		closeDb:           close,
	}

}
func (l Service) SaveProfileUser(user sql2.MotionUser, roles []sql2.RoleEnum) (sql2.Profile, error) {
	var profile sql2.Profile

	profile.Name = user.Name

	for _, role := range roles {
		field, err := l.roleRepository.FindByField("name", role)
		if err != nil {
			return sql2.Profile{}, err
		}
		profile.Roles = []sql2.Role{field}
	}

	profile.Status = sql2.ACTIVE
	profile.Birthday = user.Birthday
	profile.FamilyName = user.LastName
	profile.CreatedAt = time.Now()
	profile.User = user
	profile.UserId = user.GetId().(uint64)

	save, err := l.profileRepository.Save(profile)
	if err != nil {
		return sql2.Profile{}, err
	}
	return save, nil
}

func (l Service) FindProfileByUserId(id uint64) (sql2.Profile, error) {
	byId, err := l.profileRepository.FindWithPreloads("Roles", id)
	if err != nil {
		return sql2.Profile{}, err
	}
	return byId, nil
}
