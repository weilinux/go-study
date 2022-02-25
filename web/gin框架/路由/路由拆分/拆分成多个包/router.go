package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/web/gin/route/路由拆分/拆分成多个包/routers"
	"log"
)

// SetupRouter 封装路由
func SetupRouter() {
	r := gin.Default()

	//加载路由
	routers.LoadBlog(r)
	routers.LoadShop(r)

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
