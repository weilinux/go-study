package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//日志格式中间件
func logMiddleware(param gin.LogFormatterParams) string {
	// 自定义日志格式
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func main() {
	//创建记录日志的文件
	if f, err := os.Create("gin.log"); err != nil {
		log.Fatalln(err)
		return
	} else {
		//创建文件写入器
		gin.DefaultWriter = io.MultiWriter(f)
		// 如果需要同时将日志写入文件和控制台，请使用以下代码。
		//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	//创建路由
	r := gin.New()

	//注册全局路由
	r.Use(gin.LoggerWithFormatter(logMiddleware), gin.Recovery())

	//get请求
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"messages": "ok"})
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
