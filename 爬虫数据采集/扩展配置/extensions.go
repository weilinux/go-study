package main

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	//TODO 使用gocolly/colly操作
	// go get github.com/gocolly/colly
	// 文档：http://go-colly.org

	// 创建采集器
	c := colly.NewCollector()

	// gocolly/colly/extensions下提供三个扩展方式：

	// 根据每个请求生成随机浏览器用户代理
	extensions.RandomUserAgent(c)

	// 筛选出URL长度大于URLLengthLimit的请求
	extensions.URLLengthFilter(c, 1024)

	// 为请求设置有效的Referer HTTP标头
	extensions.Referer(c)
}
