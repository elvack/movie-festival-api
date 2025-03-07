package admin

import (
	adminModel "github.com/elvack/movie-festival-api/model/admin"
	adminRepo "github.com/elvack/movie-festival-api/repository/admin"
)

type (
	IService interface {
		Authorize(token *string) (adminId uint32, statusCode int, err error)
		Seed(req *adminModel.SeedReq) (err error)
		SignIn(reqBody *adminModel.SignInReqBody) (resData adminModel.SignInResData, statusCode int, err error)
		SignOut(adminId uint32) (err error)
	}

	service struct {
		adminRepo adminRepo.IRepo
	}
)

func NewService(adminRepo adminRepo.IRepo) IService {
	return &service{adminRepo: adminRepo}
}
