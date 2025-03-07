package movie

import (
	movieRepo "github.com/elvack/movie-festival-api/repository/movie"
	"github.com/elvack/movie-festival-api/service/movie"
	"gorm.io/gorm"
)

type controller struct {
	movieService movie.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{movieService: movie.NewService(movieRepo.NewRepo(db))}
}
