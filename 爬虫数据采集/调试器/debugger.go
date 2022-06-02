package main

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"log"
	"os"
)

func main() {
	//TODO 使用gocolly/colly操作
	// go get github.com/gocolly/colly
	// 文档：http://go-colly.org

	// 创建采集器
	c := colly.NewCollector()

	// 初始化收集器的私有变量并设置收集器的默认配置
	c.Init()

	// 收集器内的事件操作
	event := &debug.Event{
		Type:        "",  // 事件类型
		RequestID:   0,   // 标识事件的HTTP请求
		CollectorID: 0,   // 标识事件的收集器
		Values:      nil, // 值 值包含事件的键值对。不同类型的事件可以返回不同的键值对
	}

	// 创建日志调试器
	logDebugger := &debug.LogDebugger{
		Output: os.Stderr, // 指定日志输出入口  默认为os.Stderr（标准错误）
		Prefix: "error",   // 定义日志前缀（显示在每个生成的日志行的开头）
		Flag:   1,         // 定义日志属性
	}
	// 初始化日志调试器
	if err := logDebugger.Init(); err != nil {
		log.Fatalln(err)
	}
	// 接收收集器事件并将其打印到STDERR
	logDebugger.Event(event)

	// 创建web调试器
	webDebugger := &debug.WebDebugger{
		Address:         "127.0.0.1:7676", // web服务器的地址。默认值为127.0.0.1:7676
		CurrentRequests: nil,              // 当前请求数
		RequestLog:      nil,              // 请求日志
	}
	// 初始化web调试器
	if err := webDebugger.Init(); err != nil {
		log.Fatalln(err)
	}
	// 接收及机器并更新请求相关信息
	webDebugger.Event(event)

	// 将调试器附加到收集器
	c.SetDebugger(logDebugger)
}
