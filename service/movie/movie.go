package movie

import (
	movieModel "github.com/elvack/movie-festival-api/model/movie"
	movieRepo "github.com/elvack/movie-festival-api/repository/movie"
)

type (
	IService interface {
		Create(reqBody *movieModel.CreateReqBody) (err error)
		List(reqQuery *movieModel.ListReqQuery) (resData []movieModel.ListResData, count int64, err error)
		Update(req *movieModel.UpdateReq) (statusCode int, err error)
	}

	service struct {
		movieRepo movieRepo.IRepo
	}
)

func NewService(movieRepo movieRepo.IRepo) IService {
	return &service{movieRepo: movieRepo}
}
