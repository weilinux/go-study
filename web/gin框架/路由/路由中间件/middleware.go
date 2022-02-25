package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// Middleware1 中间件函数1
func Middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware1")

		//计时
		start := time.Now()

		//设置变量到上下文中 所有使用该上下文的函数都可获取此值
		c.Set("request", "client_request")

		//调用后续的处理函数 执行完成后会继续执行该行以下程序
		c.Next()

		//阻止调用后续的处理函数 执行完改函数后直接结束
		//c.Abort()

		//耗时
		elapsed := time.Since(start)
		fmt.Println("函数执行完成耗时：", elapsed)

		fmt.Println("after middleware1")
	}
}

// Middleware2 中间件函数2
func Middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware2")

		//调用后续的处理函数 执行完成后会继续执行该行以下程序
		c.Next()

		//阻止调用后续的处理函数 执行完改函数后直接结束
		//c.Abort()

		fmt.Println("after middleware2")
	}
}

// Middleware3 中间件函数3
func Middleware3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware3")

		//调用后续的处理函数 执行完成后会继续执行该行以下程序
		c.Next()

		//阻止调用后续的处理函数 执行完改函数后直接结束
		//c.Abort()

		fmt.Println("after middleware3")
	}
}

func main() {
	//创建路由
	r := gin.Default()

	//注册全局中间件
	r.Use(Middleware1(), Middleware2())

	//GET请求
	r.GET("", func(c *gin.Context) {
		//获取上下文中获取变量
		value, ok := c.Get("request")
		if !ok {
			log.Fatalln("context get error")
			return
		}
		fmt.Println("context data：", value)

		c.JSON(http.StatusOK, gin.H{"message": value})
	})

	//GET请求 注册局部中间件
	r.GET("list", Middleware3(), func(c *gin.Context) {
		//获取上下文中获取变量
		value, ok := c.Get("request")
		if !ok {
			log.Fatalln("context get error")
			return
		}
		fmt.Println("context data：", value)

		c.JSON(http.StatusOK, gin.H{"message": value})
	})

	//路由组 注册局部中间件
	userGroup := r.Group("user").Use(Middleware3())
	{
		//首页
		userGroup.GET("find", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "index"})
		})
	}

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
