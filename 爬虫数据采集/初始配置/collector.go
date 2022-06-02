package main

import (
	"context"
	"github.com/gocolly/colly"
)

func main() {
	//TODO 使用gocolly/colly操作
	// go get github.com/gocolly/colly
	// 文档：http://go-colly.org

	// 创建采集器
	c := colly.NewCollector(
		/**
		浏览器的UA字串的标准格式：浏览器标识 (操作系统标识; 加密等级标识; 浏览器语言) 渲染引擎标识版本信息
		浏览器标识	操作系统标识	加密等级标识	浏览器语言	渲染引擎	版本信息
		Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36
		*/
		//colly.UserAgent(""), 设置收集器使用的用户代理

		colly.MaxDepth(2), // 限制访问URL的递归深度

		//colly.AllowedDomains("xxx1.com", "xxx2/com"), 设置收集器使用的域白名单

		colly.ParseHTTPErrorResponse(), // 允许解析带有HTTP错误的响应

		//colly.DisallowedDomains("xxx3.com", "xxx4.com"), 设置收集器使用的域黑名单
		/**
		设置url的正则表达式
		var rg, _ = regexp.Compile(`xxx5.com/.*.html`)
		*/
		//colly.DisallowedURLFilters(rg), 设置限制访问URL的正则表达式列表。如果任何规则与URL匹配，则请求将停止

		//colly.URLFilters(rg), 设置限制访问URL的正则表达式列表。如果任何规则与URL匹配，则不会停止请求

		colly.AllowURLRevisit(), // 指示收集器允许多次下载同一URL

		colly.MaxBodySize(10*1024*1024), // 以字节为单位设置检索到的响应正文的限制

		//colly.CacheDir("./xxx1_cache"), 指定GET请求作为文件缓存的位置

		colly.IgnoreRobotsTxt(), // 指示收集器忽略目标主机的机器人设置的任何限制

		colly.ID(001), // 设置收集器的唯一标识符

		colly.Async(true), // 打开异步网络请求（启动协程）

		colly.DetectCharset(), // 支持对非utf8响应体进行字符编码检测，而无需显式的字符集声明

		/**
		创建日志调试器（打印到标准错误）
		var debugger = &debug.LogDebugger{}
		初始化日志调试器
		if err := debugger.Init(); err != nil {
			log.Fatalln(err)
		}
		*/
		//colly.Debugger(debugger), 设置收集器使用的调试器
	)

	// 初始化收集器的私有变量并设置收集器的默认配置
	c.Init()

	// 使用框架上下文（*gin.Context）替换采集器的后端http
	c.Appengine(context.Background())
}
