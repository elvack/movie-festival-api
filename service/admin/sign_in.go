package admin

import (
	"errors"
	"net/http"
	"time"

	"github.com/elvack/movie-festival-api/common"
	"github.com/elvack/movie-festival-api/library/encrypt"
	adminModel "github.com/elvack/movie-festival-api/model/admin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func (s *service) SignIn(reqBody *adminModel.SignInReqBody) (resData adminModel.SignInResData, statusCode int, err error) {
	admin, err := s.adminRepo.Take([]string{"encrypted_password", "id"}, &adminModel.Admin{Email: reqBody.Email})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resData, http.StatusBadRequest, errors.New(common.EmailOrPasswordIsIncorrect)
		}
		return resData, http.StatusInternalServerError, err
	}
	if err = encrypt.CompareHashAndPassword(&admin.EncryptedPassword, &reqBody.Password); err != nil {
		return resData, http.StatusBadRequest, errors.New(common.EmailOrPasswordIsIncorrect)
	}
	authClaims := adminModel.AuthClaims{
		AdminId: admin.Id,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)),
		},
	}
	resData.Token, err = encrypt.NewTokenWithClaims(authClaims)
	if err != nil {
		return resData, http.StatusInternalServerError, err
	}
	if err = s.adminRepo.Update(&admin.Id, &map[string]any{"token": resData.Token}); err != nil {
		return resData, http.StatusInternalServerError, err
	}
	return resData, http.StatusOK, nil
}
