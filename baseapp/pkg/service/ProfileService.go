package service

import (
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type ProfileService struct {
	profileRepository *repository.MotionSQLRepository[sqlDomain.Profile]
	roleRepository    *repository.MotionSQLRepository[sqlDomain.Role]
}

func NewProfileService(conections *db.Connections) *ProfileService {
	return &ProfileService{
		profileRepository: repository.NewProfileRepository(conections.GetPgsqTemplate()),
		roleRepository:    repository.NewRoleRepository(conections.GetPgsqTemplate()),
	}

}
func (l *ProfileService) SaveProfileUser(user sqlDomain.MotionUser, roles []sqlDomain.RoleEnum) (sqlDomain.Profile, error) {
	var profile sqlDomain.Profile

	profile.Name = user.Name

	for _, role := range roles {
		field, err := l.roleRepository.FindByField("name", role)
		if err != nil {
			return sqlDomain.Profile{}, err
		}
		profile.Roles = append(profile.Roles, field)
	}

	profile.Status = domain.ACTIVE
	profile.Birthday = user.Birthday
	profile.FamilyName = user.LastName
	profile.CreatedAt = time.Now()
	profile.User = user

	save, err := l.profileRepository.Save(profile)
	if err != nil {
		return sqlDomain.Profile{}, err
	}
	return save, nil
}

func (l *ProfileService) FindProfileByUserId(id uint64) (sqlDomain.Profile, error) {
	byId, err := l.profileRepository.FindWithPreloads("Roles", id)
	if err != nil {
		return sqlDomain.Profile{}, err
	}
	return byId, nil
}
