package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Errno int

const (
	ok        = 200 // 成功
	BindError = 201 // 参数解析错误
)

var HTTPErrno = map[Errno]string{
	BindError: "参数解析错误",
}

func NewError(errno Errno, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    BindError,
		"message": HTTPErrno[BindError],
	})
}