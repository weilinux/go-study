package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/web/gin/route/路由拆分/拆分到不同的APP/app/blog"
	"github.com/my/repo/web/gin/route/路由拆分/拆分到不同的APP/app/shop"

	"log"
)

//路由配置结构体
type option func(engine *gin.Engine)

//路由配置切片
var options []option

//定义路由的数据类型
var r *gin.Engine

// Include 注册app的路由配置
func include(o ...option) {
	options = append(options, o...)
}

// 初始化
func init() {
	//将路由配置加载到切片中
	include(blog.Router, shop.Router)

	//创建不带中间件的路由
	r = gin.New()

	//循环路由配置切片
	for _, o := range options {
		//将路由配置挂载到路由上
		o(r)
	}
}

// SetupRouter 封装路由
func SetupRouter() {
	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
