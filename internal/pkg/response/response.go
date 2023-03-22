package response

import (
	cpdserr "cpds/cpds-analyzer/internal/pkg/errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleOK(ctx *gin.Context, data interface{}) {
	r := &ResponseBody{
		Status:    http.StatusOK,
		Code:      cpdserr.SUCCESS,
		Message:   "Success",
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
	ctx.JSON(http.StatusOK, r)
}

func HandleError(ctx *gin.Context, httpStatus int, err error) {
	r := &ResponseBody{
		Status:    httpStatus,
		Code:      int(err.(*cpdserr.Error).ResultCode),
		Message:   err.Error(),
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
	ctx.Error(err)
	ctx.AbortWithStatusJSON(httpStatus, r)
}
