package admin

import (
	"errors"

	"github.com/elvack/movie-festival-api/common"
	"github.com/elvack/movie-festival-api/library/encrypt"
	adminModel "github.com/elvack/movie-festival-api/model/admin"
	"gorm.io/gorm"
)

func (s *service) Seed(req *adminModel.SeedReq) (err error) {
	if _, err = s.adminRepo.Take([]string{"id"}, &adminModel.Admin{Email: req.Email}); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
	} else {
		return errors.New(common.EmailAlreadyExists)
	}
	if err = encrypt.GenerateFromPassword(&req.Password); err != nil {
		return
	}
	admin := adminModel.Admin{
		Email: req.Email,
		EncryptedPassword: req.Password,
	}
	return s.adminRepo.Create(&admin)
}
