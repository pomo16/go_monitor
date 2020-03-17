package test

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/utils"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	jwt := utils.NewJWT()
	token, _ := jwt.CreateToken(utils.CustomClaims{
		UserName: "",
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                   // 签名生效时间
			ExpiresAt: time.Now().Add(consts.TokenExpired).Unix(), // 过期时间15分钟
			Issuer:    "pomo",                                     //签名的发行者
		},
	})
	fmt.Println(token)

	claims, err := utils.NewJWT().ParseToken(token)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Printf("claims: %v", claims.UserName)
}

func TestMD5(t *testing.T) {
	fmt.Println(utils.Md5AddSalt("pomo", consts.UserNameSalt, true))
	fmt.Println(utils.Md5AddSalt("123456", consts.PasswordSalt, false))
}
