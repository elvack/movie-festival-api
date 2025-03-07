package admin

import (
	"net/http"

	"github.com/elvack/movie-festival-api/common"
	"github.com/elvack/movie-festival-api/library/response"
	adminModel "github.com/elvack/movie-festival-api/model/admin"
	"github.com/gin-gonic/gin"
)

func (c *controller) SignIn(ctx *gin.Context) {
	var reqBody adminModel.SignInReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resData, statusCode, err := c.adminService.SignIn(&reqBody)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, statusCode, common.SignedInSuccessfully, resData)
}
