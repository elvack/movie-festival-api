package movie

import (
	"net/http"

	"github.com/elvack/movie-festival-api/common"
	"github.com/elvack/movie-festival-api/library/response"
	movieModel "github.com/elvack/movie-festival-api/model/movie"
	"github.com/gin-gonic/gin"
)

func (c *controller) Create(ctx *gin.Context) {
	var reqBody movieModel.CreateReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := c.movieService.Create(&reqBody); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.CreatedSuccessfully, nil)
}
