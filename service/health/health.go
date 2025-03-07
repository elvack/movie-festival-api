package health

import healthRepo "github.com/elvack/movie-festival-api/repository/health"

type (
	IService interface {
		Check() (err error)
	}

	service struct {
		healthRepo healthRepo.IRepo
	}
)

func NewService(healthRepo healthRepo.IRepo) IService {
	return &service{healthRepo: healthRepo}
}
