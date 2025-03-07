package movie

import movieModel "github.com/elvack/movie-festival-api/model/movie"

func (r *repo) Create(movie *movieModel.Movie) (err error) {
	return r.db.Create(movie).Error
}
