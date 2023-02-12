package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/common"
	"runtime/debug"
)

func CatchError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic err: %v\n, trace: %v\n", err, string(debug.Stack()))
			common.Error(c, common.NOT_CATCH_ERROR, nil)
			c.Abort()
		}
	}()
	c.Next()
}
