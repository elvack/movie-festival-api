package router

import (
	"net/http"

	"github.com/elvack/movie-festival-api/library/response"
	adminModel "github.com/elvack/movie-festival-api/model/admin"
	adminRepo "github.com/elvack/movie-festival-api/repository/admin"
	"github.com/elvack/movie-festival-api/service/admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func authorize(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqHeader adminModel.ReqHeader
		if err := ctx.ShouldBindHeader(&reqHeader); err != nil {
			response.Error(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		adminService := admin.NewService(adminRepo.NewRepo(db))
		adminId, statusCode, err := adminService.Authorize(&reqHeader.Authorization)
		if err != nil {
			response.Error(ctx, statusCode, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("adminId", adminId)
	}
}
