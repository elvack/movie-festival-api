package health

import (
	"database/sql"

	healthRepo "github.com/elvack/movie-festival-api/repository/health"
	"github.com/elvack/movie-festival-api/service/health"
)

type controller struct {
	healthService health.IService
}

func NewController(db *sql.DB) *controller {
	return &controller{healthService: health.NewService(healthRepo.NewRepo(db))}
}
