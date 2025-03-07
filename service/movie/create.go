package movie

import movieModel "github.com/elvack/movie-festival-api/model/movie"

func (s *service) Create(reqBody *movieModel.CreateReqBody) (err error) {
	movie := movieModel.Movie{
		Description: reqBody.Description,
		Duration: reqBody.Duration,
		Title: reqBody.Title,
		WatchUrl: reqBody.WatchUrl,
	}
	return s.movieRepo.Create(&movie)
}
