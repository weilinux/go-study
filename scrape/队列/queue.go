package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"log"
)

func main() {
	//TODO 使用gocolly/colly操作
	// go get github.com/gocolly/colly
	// 文档：http://go-colly.org

	// 创建采集器
	c := colly.NewCollector()

	// 初始化收集器的私有变量并设置收集器的默认配置
	c.Init()

	// 已默认存储方式创建队列	threads：线程数
	q, err := queue.New(10, &queue.InMemoryQueueStorage{
		MaxSize: 1024, //队列容量
	})
	if err != nil {
		log.Fatalln(err)
	}

	// 判断队列是否为空
	fmt.Println("队列是否为空：", q.IsEmpty())

	// 获取队列大小
	if size, err := q.Size(); err == nil {
		fmt.Println("队列大小：", size)
	} else {
		log.Fatalln(err)
	}

	// 向队列添加新请求
	if err = q.AddRequest(&colly.Request{
		URL:                       nil, // 地址
		Headers:                   nil, // 请求头
		Ctx:                       nil, // 上下文
		Depth:                     0,   // 请求的父级数
		Method:                    "",  // 请求的http方法
		Body:                      nil, // 用于POST/PUT请求的请求正文
		ResponseCharacterEncoding: "",  // 响应正文的字符编码
		ID:                        0,   // 请求的唯一标识符
		ProxyURL:                  "",  // 处理请求的代理地址
	}); err != nil {
		log.Fatalln(err)
	}

	// 将新URL添加到队列
	if err = q.AddURL("www.com/q/s"); err != nil {
		log.Fatalln(err)
	}

	// 启动线程并调用收集器来执行请求
	if err = q.Run(c); err != nil {
		log.Fatalln(err)
	}

}
