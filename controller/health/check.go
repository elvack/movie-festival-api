package health

import (
	"net/http"

	"github.com/elvack/movie-festival-api/common"
	"github.com/elvack/movie-festival-api/library/response"
	"github.com/gin-gonic/gin"
)

func (c *controller) Check(ctx *gin.Context) {
	if err := c.healthService.Check(); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.CheckedSuccessfully, nil)
}
