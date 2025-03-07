package response

import "github.com/gin-gonic/gin"

func SuccessWithPagination(ctx *gin.Context, statusCode int, message string, data any, current, size int, total int64) {
	ctx.JSON(statusCode, resWithPage{
		Data: data,
		Message: message,
		Page: page{
			Current: current,
			Size: size,
			Total: total,
		},
	})
}
