package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"project/common"
	"project/global"
)

// NeedLogin 仅获取UserId，可能会获取不到
func NeedLogin(c *gin.Context) {

	userId := ""
	tokenStr, _ := c.Cookie("ATK")

	if tokenStr != "" {
		if token, err := parseIdFromToken(tokenStr); token != nil && err == nil {

			fmt.Println(token)
			claims := token.Claims.(*common.CustomClaims)
			if claims.Id != "" {
				userId = claims.Id
			}
		}
	}
	c.Set("userId", userId)
}

// CheckLogin 校验Token
func CheckLogin(c *gin.Context) {

	userId := ""
	tokenStr, _ := c.Cookie("ATK")

	if tokenStr == "" {
		common.Error(c, common.NEED_LOGIN, nil)
		c.Abort()
		return
	}
	token, err := parseIdFromToken(tokenStr)

	if token == nil || err != nil {
		common.Error(c, common.NEED_LOGIN, nil)
		c.Abort()
		return
	}

	claims := token.Claims.(*common.CustomClaims)
	if claims.Id != "" {
		userId = claims.Id
	}
	c.Set("userId", userId)
}

func parseIdFromToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &common.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return jwt.ErrInvalidKey, nil
		}
		return []byte(global.App.Config.Jwt.Secret), nil
	})
	return token, err
}
