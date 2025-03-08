package movie

import (
	"net/http"

	"github.com/elvack/movie-festival-api/common"
	"github.com/elvack/movie-festival-api/library/response"
	movieModel "github.com/elvack/movie-festival-api/model/movie"
	"github.com/gin-gonic/gin"
)

func (c *controller) Update(ctx *gin.Context) {
	var req movieModel.UpdateReq
	if err := ctx.ShouldBindUri(&req.Path); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctx.ShouldBindJSON(&req.Body); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if statusCode, err := c.movieService.Update(&req); err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.UpdatedSuccessfully, nil)
}
