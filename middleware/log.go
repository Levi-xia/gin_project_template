package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/common"
	"project/global"
	"strconv"
	"time"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func WriteActionLog(c *gin.Context) {

	start := time.Now()
	blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	c.Next()

	traceId := c.GetString("traceId")

	var (
		userId string
		input  string
		output string
	)

	// 用户Id
	if v, exists := c.Keys["userId"]; exists {
		userId = fmt.Sprintf("%v", v)
	}

	// Uri
	uri := c.Request.URL.String()

	// 输入参数
	if inputList, err := common.GetRequestInputs(c); err == nil {
		if inputStr, err := json.Marshal(inputList); err == nil {
			input = string(inputStr)
		}
	}

	// 输出参数
	respList := common.Response{}

	json.Unmarshal(blw.body.Bytes(), &respList)

	outputStr, err := json.Marshal(respList)

	if err != nil {
		panic(err)
	}
	output = string(outputStr)

	// 接口耗时
	cost := time.Since(start).String()

	// 打印日志
	if respList.ErrorNo > 0 {
		global.App.Log.Warn("",
			zap.String("errno", strconv.Itoa(respList.ErrorNo)),
			zap.String("userId", userId),
			zap.String("traceId", traceId),
			zap.String("uri", uri),
			zap.String("input", input),
			zap.String("output", output),
			zap.String("cost", cost))
	} else {
		global.App.Log.Info("",
			zap.String("errno", strconv.Itoa(respList.ErrorNo)),
			zap.String("userId", userId),
			zap.String("traceId", traceId),
			zap.String("uri", uri),
			zap.String("input", input),
			zap.String("output", output),
			zap.String("cost", cost))
	}

}
