package common

import (
	"github.com/dgrijalva/jwt-go"
	"project/global"
	"time"
)

// CustomClaims 自定义 Claims
type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType       = "bearer"
	AppGuardName    = "app"
	WeChatGuardName = "wxapp"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func CreateToken(userId string, guardName string) (TokenOutPut, *jwt.Token, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.JwtTtl,
				Id:        userId,
				Issuer:    guardName, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)
	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))

	tokenData := TokenOutPut{
		tokenStr,
		int(global.App.Config.Jwt.JwtTtl),
		TokenType,
	}
	return tokenData, token, err
}
