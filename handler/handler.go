package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/pkg/errno"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func NotFound(c *gin.Context) {
	SendResponse(c, errno.ErrNotFound, nil)
}
