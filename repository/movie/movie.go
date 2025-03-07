package movie

import (
	movieModel "github.com/elvack/movie-festival-api/model/movie"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(movie *movieModel.Movie) (err error)
		List(reqQuery *movieModel.ListReqQuery) (resData []movieModel.ListResData, count int64, err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
