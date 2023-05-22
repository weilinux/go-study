package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	//创建路由
	r := gin.Default()

	//异步
	r.GET("/async", func(c *gin.Context) {
		//获取上下文的只读副本
		context := c.Copy()

		//异步处理 在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Println("异步执行：", context.Request.URL.Path)
		}()
	})

	//同步
	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		fmt.Println("同步执行：", c.Request.URL.Path)
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
