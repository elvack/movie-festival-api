package admin

import (
	"errors"
	"net/http"

	"github.com/elvack/movie-festival-api/common"
	"github.com/elvack/movie-festival-api/library/encrypt"
	adminModel "github.com/elvack/movie-festival-api/model/admin"
	"gorm.io/gorm"
)

func (s *service) Authorize(token *string) (adminId uint32, statusCode int, err error) {
	tokenRaw, claims, err := encrypt.Parse(*token)
	if err != nil {
		return adminId, http.StatusUnauthorized, err
	}
	adminId = uint32(claims["AdminId"].(float64))
	admin, err := s.adminRepo.Take([]string{"token"}, &adminModel.Admin{Id: adminId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return adminId, http.StatusUnauthorized, errors.New(common.AdminNotFound)
		}
		return adminId, http.StatusInternalServerError, err
	}
	if tokenRaw != admin.Token {
		return adminId, http.StatusUnauthorized, errors.New(common.AdminHasSignedOut)
	}
	return adminId, http.StatusOK, nil
}
