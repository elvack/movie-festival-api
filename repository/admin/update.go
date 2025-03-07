package admin

import adminModel "github.com/elvack/movie-festival-api/model/admin"

func (r *repo) Update(id *uint32, values *map[string]any) (err error) {
	return r.db.Model(adminModel.Admin{Id: *id}).Updates(values).Error
}
