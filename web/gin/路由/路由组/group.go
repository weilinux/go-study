package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//创建路由
	r := gin.Default()

	//分组
	videoGroup := r.Group("video")
	{
		//首页
		videoGroup.GET("index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "index"})
		})

		//上传
		videoGroup.POST("upload", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "upload"})
		})

		//修改
		videoGroup.PUT("update", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "update"})
		})

		//删除
		videoGroup.DELETE("remove", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "remove"})
		})
	}

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
