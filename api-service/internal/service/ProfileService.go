package service

import (
	sqlDomain2 "github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/src/emailSender"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/repository"
)

type ProfileService struct {
	profileRepository *repository.MotionSQLRepository[sqlDomain2.Profile]
	roleRepository    *repository.MotionSQLRepository[sqlDomain2.Role]
}

func NewProfileService(conections *db.Connections) *ProfileService {
	return &ProfileService{
		profileRepository: repository.NewProfileRepository(conections.GetPgsqTemplate()),
		roleRepository:    repository.NewRoleRepository(conections.GetPgsqTemplate()),
	}
}

func (l *ProfileService) SaveProfileUser(user sqlDomain2.MotionUser, roles []sqlDomain2.RoleEnum) (sqlDomain2.Profile, error) {
	var profile sqlDomain2.Profile

	profile.Name = user.Name

	for _, role := range roles {
		field, err := l.roleRepository.FindByField("name", role)
		if err != nil {
			return sqlDomain2.Profile{}, err
		}
		profile.Roles = append(profile.Roles, field)
	}

	profile.Status = sqlDomain2.EMAIL_SYNC
	profile.CreatedAt = time.Now()
	profile.MotionUserId = user.Id
	code := emailSender.GenerateEmailCode()
	profile.Code = code

	save, err := l.profileRepository.Save(profile)
	if err != nil {
		return sqlDomain2.Profile{}, err
	}
	return save, nil
}

func (l *ProfileService) FindProfileByUserId(id uint64) (sqlDomain2.Profile, error) {
	byId, err := l.profileRepository.FindWithPreloads("Roles", id)
	if err != nil {
		return sqlDomain2.Profile{}, err
	}
	return byId, nil
}
