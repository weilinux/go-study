package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//创建路由
	r := gin.Default()

	//外部重定向 通过Redirect跳转到外部页面
	r.GET("index", func(c *gin.Context) {
		//重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	//内部重定向 通过c.Request.URL.Path 设置跳转的指定的路径
	r.GET("test1", func(c *gin.Context) {
		//设置请求路径 这里必须要加斜线
		c.Request.URL.Path = "/test2"

		//执行context
		r.HandleContext(c)
	})

	r.GET("test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
