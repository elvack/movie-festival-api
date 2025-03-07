package admin

import (
	adminModel "github.com/elvack/movie-festival-api/model/admin"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(admin *adminModel.Admin) (err error)
		Take(selectParams []string, conditions *adminModel.Admin) (admin adminModel.Admin, err error)
		Update(id *uint32, values *map[string]any) (err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
