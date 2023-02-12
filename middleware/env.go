package middleware

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

func InitEnv(c *gin.Context) {

	// 写入全局TraceId
	c.Set("traceId", strconv.Itoa(rand.Int()))
}
