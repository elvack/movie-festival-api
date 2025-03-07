package movie

import movieModel "github.com/elvack/movie-festival-api/model/movie"

func (s *service) List(reqQuery *movieModel.ListReqQuery) (resData []movieModel.ListResData, count int64, err error) {
	return s.movieRepo.List(reqQuery)
}
