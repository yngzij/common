package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": ok,
	})
}

func ListData(data interface{}, total int64, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":  ok,
		"total": total,
		"data":  data,
	})
}

func OneData(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": ok,
		"data": data,
	})
}
