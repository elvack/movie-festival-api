package movie

import movieModel "github.com/elvack/movie-festival-api/model/movie"

func (r *repo) Update(id *uint32, values *map[string]any) (err error) {
	return r.db.Model(movieModel.Movie{Id: *id}).Updates(values).Error
}
