package movie

import (
	"errors"
	"net/http"

	"github.com/elvack/movie-festival-api/common"
	movieModel "github.com/elvack/movie-festival-api/model/movie"
	"gorm.io/gorm"
)

func (s *service) Update(req *movieModel.UpdateReq) (statusCode int, err error) {
	if _, err = s.movieRepo.Take([]string{"id"}, &movieModel.Movie{Id: req.Path.Id}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, errors.New(common.MovieNotFound)
		}
		return http.StatusInternalServerError, err
	}
	values := map[string]any{
		"description": req.Body.Description,
		"duration": req.Body.Duration,
		"title": req.Body.Title,
		"watch_url": req.Body.WatchUrl,
	}
	if err = s.movieRepo.Update(&req.Path.Id, &values); err != nil {
		return http.StatusInternalServerError, err
	}
	return
}
