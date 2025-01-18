package common

import "github.com/gin-gonic/gin"

func ReturnError(ctx *gin.Context, status int, body any, err error) {
	ctx.JSON(status, body)
}
