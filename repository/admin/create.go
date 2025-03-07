package admin

import adminModel "github.com/elvack/movie-festival-api/model/admin"

func (r *repo) Create(admin *adminModel.Admin) (err error) {
	return r.db.Create(admin).Error
}
