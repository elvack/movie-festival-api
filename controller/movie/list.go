package movie

import (
	"net/http"

	"github.com/elvack/movie-festival-api/library/pagination"
	"github.com/elvack/movie-festival-api/library/response"
	movieModel "github.com/elvack/movie-festival-api/model/movie"
	"github.com/gin-gonic/gin"
)

func (c *controller) List(ctx *gin.Context) {
	var reqQuery movieModel.ListReqQuery
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	reqQuery.Offset = pagination.Offset(&reqQuery.Limit, &reqQuery.Page)
	resData, count, err := c.movieService.List(&reqQuery)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessWithPagination(ctx, http.StatusOK, "ok", resData, reqQuery.Page, reqQuery.Limit, count)
}
