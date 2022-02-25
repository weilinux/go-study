package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//创建路由
	r := gin.Default()

	//get请求
	r.GET("cookie", func(c *gin.Context) {
		//获取cookie
		cookie, err := c.Cookie("clientCookie")

		//判断客户端是否携带cookie
		if err != nil {
			cookie = "NotFound"

			//设置cookie
			/**
			name     string  名称
			value    string  值
			maxAge   int     有效时长，单位秒
			path     string  cookie所在目录
			domain   string  cookie所在域 （作用范围）
			secure   bool    是否只能通过https访问
			httpOnly bool    是否可以通过js代码进行操作
			*/
			c.SetCookie("clientCookie", "clientValue", 3600, "/", "127.0.0.1", false, true)
		}

		fmt.Println("cookie的值是：", cookie)
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
