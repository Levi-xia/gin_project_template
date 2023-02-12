package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ErrorNo int    `json:"errno"`
	ErrMsg  string `json:"errmsg"`
	Data    any    `json:"data"`
}

type Options struct {
	ShowDetail  bool   `json:"showDetail"`
	FullMessage string `json:"fullMessage""`
	LogLevel    string `json:"LogLevel"`
}

func Success(c *gin.Context, data any) {
	resp := &Response{
		ErrorNo: 0,
		ErrMsg:  "ok",
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func Error(c *gin.Context, errCode int, data any) {

	errMsg := "unknown error"

	if v, ok := AppCodes[errCode]; ok {
		errMsg = v
	}
	resp := &Response{
		ErrorNo: errCode,
		ErrMsg:  errMsg,
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func WithError(c *gin.Context, errCode int, data any, options Options) {

	errMsg := "unknown error"

	if !options.ShowDetail {
		if v, ok := AppCodes[errCode]; ok {
			errMsg = v
		}
	} else if options.ShowDetail && options.FullMessage != "" {
		errMsg = options.FullMessage
	}
	resp := &Response{
		ErrorNo: errCode,
		ErrMsg:  errMsg,
		Data:    data,
	}

	c.JSON(http.StatusOK, resp)
}