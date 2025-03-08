package movie

import movieModel "github.com/elvack/movie-festival-api/model/movie"

func (r *repo) Take(selectParams []string, conditions *movieModel.Movie) (movie movieModel.Movie, err error) {
	return movie, r.db.Select(selectParams).Take(&movie, conditions).Error
}
