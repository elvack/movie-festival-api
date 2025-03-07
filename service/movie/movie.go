package movie

import (
	movieModel "github.com/elvack/movie-festival-api/model/movie"
	movieRepo "github.com/elvack/movie-festival-api/repository/movie"
)

type (
	IService interface {
		Create(reqBody *movieModel.CreateReqBody) (err error)
	}

	service struct {
		movieRepo movieRepo.IRepo
	}
)

func NewService(movieRepo movieRepo.IRepo) IService {
	return &service{movieRepo: movieRepo}
}
