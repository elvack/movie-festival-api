package response

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, statusCode int, message string, data any) {
	ctx.JSON(statusCode, response{
		Data: data,
		Message: message,
	})
}
