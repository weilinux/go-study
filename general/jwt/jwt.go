package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// jwt令牌
// go get github.com/dgrijalva/jwt-go

// jwt密码盐
var secret = []byte("secret1265")

// Claims 用户信息结构体
type Claims struct {
	UserId   uint   `form:"userId" json:"userId"`
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	jwt.StandardClaims
}

// GenToken 生成jwt
func GenToken(claims *Claims) (string, error) {
	//当前时间
	now := time.Now()
	//目标接受者
	claims.Audience = "xxx"
	//过期时间
	claims.ExpiresAt = now.Add(2 * time.Hour).Unix()
	//签名id，用于防止jwt被重放
	claims.Id = "0001"
	//签发时间
	claims.IssuedAt = now.Unix()
	//签发人
	claims.Issuer = "127.0.0.1"
	//在...时间之后才生效
	claims.NotBefore = now.Add(1 * time.Minute).Unix()
	//签名主题（场景）
	claims.Subject = "userToken"

	//使用指定的签名方法(hash)创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secret)
}

// ParseToken 解析jwt
func ParseToken(tokenString string) (*Claims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func main() {
	//创建路由
	r := gin.Default()

	//get请求
	r.POST("token", func(c *gin.Context) {
		var claims *Claims
		err := c.ShouldBind(&claims)
		if err != nil {
			log.Fatalln(err)
			return
		}

		//生成jwt
		tokenString, err := GenToken(claims)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		fmt.Println("tokenString：", tokenString)

		//解析jwt
		tokenClaims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": err.Error()})
			return
		}
		fmt.Println("tokenClaims：", tokenClaims)
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
