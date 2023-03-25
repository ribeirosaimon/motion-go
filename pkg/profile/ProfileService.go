package profile

import (
	"database/sql"
	"time"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/repository"
	"gorm.io/gorm"
)

type Service struct {
	profileRepository repository.MotionRepository[domain.Profile]
	roleRepository    repository.MotionRepository[domain.Role]
	closeDb           *sql.DB
}

func NewProfileService(conn *gorm.DB, close *sql.DB) Service {
	return Service{
		profileRepository: repository.NewProfileRepository(conn),
		roleRepository:    repository.NewRoleRepository(conn),
		closeDb:           close,
	}

}
func (l Service) SaveProfileUser(user domain.MotionUser, roles []domain.RoleEnum) (domain.Profile, error) {
	var profile domain.Profile

	profile.Name = user.Name

	for _, role := range roles {
		field, err := l.roleRepository.FindByField("name", role)
		if err != nil {
			return domain.Profile{}, err
		}
		profile.Roles = []domain.Role{field}
	}

	profile.Status = domain.ACTIVE
	profile.Birthday = user.Birthday
	profile.FamilyName = user.LastName
	profile.CreatedAt = time.Now()
	profile.User = user
	profile.UserId = user.GetId().(uint64)

	save, err := l.profileRepository.Save(profile)
	if err != nil {
		return domain.Profile{}, err
	}
	return save, nil
}

func (l Service) FindProfileByUserId(id uint64) (domain.Profile, error) {
	byId, err := l.profileRepository.FindWithPreloads("Roles", id)
	if err != nil {
		return domain.Profile{}, err
	}
	return byId, nil
}
