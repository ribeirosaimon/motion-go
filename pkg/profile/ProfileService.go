package profile

import (
	"time"

	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/config/database"
	"github.com/ribeirosaimon/motion-go/repository"
)

type Service struct {
	profileRepository repository.MotionRepository[domain.Profile]
}

func NewProfileService() Service {
	connect, _ := database.Connect()
	return Service{
		profileRepository: repository.NewProfileRepository(connect),
	}

}
func (l Service) SaveProfileUser(user domain.MotionUser) (domain.Profile, error) {
	var profile domain.Profile

	profile.Name = user.Name

	profile.Roles = []domain.Role{
		{Name: domain.USER},
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
