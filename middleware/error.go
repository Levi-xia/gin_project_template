package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/common"
)

func CatchError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic err: %v\n", err)
			common.Error(c, common.NOT_CATCH_ERROR, nil)
			c.Abort()
		}
	}()
	c.Next()
}
