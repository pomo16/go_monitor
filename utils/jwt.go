package utils

import (
	"github.com/dgrijalva/jwt-go"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"time"
)

//JWT 签名结构
type JWT struct {
	SigningKey []byte
}

//CustomClaims 自定义载荷
type CustomClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

//NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(consts.Secret),
	}
}

//CreateToken 生成一个Token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	}
	return nil, exceptions.ErrToken
}

//RefreshToken 更新Token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(consts.TokenExpired).Unix()
		newToken, err := j.CreateToken(*claims)
		if err != nil {
			return "", exceptions.ErrToken
		}
		return newToken, nil
	}
	return "", exceptions.ErrToken
}
