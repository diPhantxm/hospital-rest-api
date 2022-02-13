package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorAbort(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"status_code": statusCode,
		"message":     message,
	})
}

func emptyObject(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func emptySlice(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []gin.H{})
}
