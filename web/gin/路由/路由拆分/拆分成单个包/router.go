package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//封装处理函数
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}

// SetupRouter 封装路由
func SetupRouter() {
	r := gin.Default()

	r.GET("", helloHandler)

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
