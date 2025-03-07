package admin

import (
	"net/http"

	"github.com/elvack/movie-festival-api/common"
	"github.com/elvack/movie-festival-api/library/response"
	"github.com/gin-gonic/gin"
)

func (c *controller) SignOut(ctx *gin.Context) {
	if err := c.adminService.SignOut(ctx.MustGet("adminId").(uint32)); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SignedOutSuccessfully, nil)
}
