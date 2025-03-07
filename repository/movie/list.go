package movie

import movieModel "github.com/elvack/movie-festival-api/model/movie"

func (r *repo) List(reqQuery *movieModel.ListReqQuery) (resData []movieModel.ListResData, count int64, err error) {
	resData = make([]movieModel.ListResData, 0)
	return resData, count, r.db.Select("created_at, description, duration, id, title, updated_at, watch_url").Model(movieModel.Movie{}).Count(&count).Limit(reqQuery.Limit).Offset(reqQuery.Offset).Scan(&resData).Error
}
