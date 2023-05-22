package main

import (
	"github.com/gocolly/colly"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// 代理配置
var proxies = []*url.URL{
	{Host: "127.0.0.1:8080"},
	{Host: "127.0.0.1:8081"},
}

func randomProxySwitcher(_ *http.Request) (*url.URL, error) {
	// 随机种子
	rand.Seed(time.Now().UnixMilli())

	// 随机切换代理
	return proxies[rand.Intn(len(proxies))], nil
}

func main() {
	//TODO 使用gocolly/colly操作
	// go get github.com/gocolly/colly
	// 文档：http://go-colly.org

	// 创建采集器
	c := colly.NewCollector()

	// 初始化收集器的私有变量并设置收集器的默认配置
	c.Init()

	// 内置的通过轮询方式实现代理切换的函数
	/*switcher, err := proxy.RoundRobinProxySwitcher(
		"socks5://127.0.0.1:1235",
		"socks5://127.0.0.1:1236",
		"http://127.0.0.1:1237",
	)
	if err != nil {
		log.Fatalln(err)
	}*/

	// 设置自定义代理切换器函数
	c.SetProxyFunc(randomProxySwitcher)
}
