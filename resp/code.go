package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Errno int

const (
	ok               = 200 // 成功
	BindError        = 201 // 参数解析错误
	InternalError    = 202
	IdNoneError      = 203
	CreateRepeated   = 204
	CreateGroupError = 205
)

var HTTPErrno = map[Errno]string{
	BindError:        "参数解析错误",
	InternalError:    "服务器内部错误",
	IdNoneError:      "没有查找到用户",
	CreateRepeated:   "重复创建",
	CreateGroupError: "创建群组失败",
}

func NewError(errno Errno, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    errno,
		"message": HTTPErrno[errno],
	})
}
