package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/redisstorage"
	"log"
	"net/url"
)

func main() {
	//TODO 使用gocolly/colly操作
	// go get github.com/gocolly/colly

	// 创建采集器
	c := colly.NewCollector()

	// 初始化收集器的私有变量并设置收集器的默认配置
	c.Init()

	// redis存储器
	// go get -u github.com/gocolly/redisstorage
	inRedisStorage := &redisstorage.Storage{
		Address:  "127.0.0.1:6379", // 地址 ip:port
		Password: "",               // 密码
		DB:       0,                // 数据库
		Prefix:   "c01",            // 前缀
	}

	// 初始化
	if err := inRedisStorage.Init(); err != nil {
		log.Fatalln(err)
	}

	// 清空
	if err := inRedisStorage.Clear(); err != nil {
		log.Fatalln(err)
	}

	// 创建请求信息
	request := &colly.Request{
		URL:                       nil, // 地址
		Headers:                   nil, // 请求头
		Ctx:                       nil, // 上下文
		Depth:                     0,   // 请求的父级数
		Method:                    "",  // 请求的http方法
		Body:                      nil, // 用于POST/PUT请求的请求正文
		ResponseCharacterEncoding: "",  // 响应正文的字符编码
		ID:                        0,   // 请求的唯一标识符
		ProxyURL:                  "",  // 处理请求的代理地址
	}
	// 将请求信息json序列化
	if jsonRequest, err := json.Marshal(request); err == nil {
		// 向队列添加新请求
		if err = inRedisStorage.AddRequest(jsonRequest); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

	// 获取go-redis客户端
	fmt.Println("go-redis客户端：", inRedisStorage.Client)

	// 从队列中获取请求
	if getRequest, err := inRedisStorage.GetRequest(); err != nil {
		// 申明请求信息变量
		request2 := &colly.Request{}
		// 将请求信息反序列化并绑定
		if err = json.Unmarshal(getRequest, request2); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

	// 访问该标识事件的HTTP请求
	if err := inRedisStorage.Visited(3); err != nil {
		log.Fatalln(err)
	}

	// 是否已访问该标识事件的HTTP请求
	if visited, err := inRedisStorage.IsVisited(4); err != nil {
		fmt.Println("是否访问：", visited)
	} else {
		log.Fatalln(err)
	}

	// 创建已解析的url信息
	u := &url.URL{
		Scheme:      "",               // url方案	示例：[scheme:][//[userinfo@]host][/]path[?query][#fragment]
		Opaque:      "",               // 编码后的不透明数据
		User:        nil,              // 用户名和密码信息
		Host:        "127.0.0.1:8080", // host或host:port
		Path:        "",               // 路径（相对路径可以省略前导斜杠）
		RawPath:     "",               // 编码路径提示
		ForceQuery:  false,            // 即使RawQuery为空，也要附加查询（“？”）
		RawQuery:    "",               // 编码后的查询字符串，没有'?'
		Fragment:    "",               // 引用的片段（文档位置），没有'#'
		RawFragment: "",               // 编码片段提示
	}

	// 设置存储器对应url的cookie
	inRedisStorage.SetCookies(u, "_ga=GA1.2.1611472128.1650815524; _gid=GA1.2.2080811677.1652022429; __atuvc=2|17,0|18,5|19")

	// 获取存储器对应url的cookie
	fmt.Println("cookie信息：", inRedisStorage.Cookies(u))

	// 获取队列大小
	if size, err := inRedisStorage.QueueSize(); err == nil {
		fmt.Println("队列大小：", size)
	} else {
		log.Fatalln(err)
	}

	// 设置采集器的redis存储器
	if err := c.SetStorage(inRedisStorage); err != nil {
		log.Fatalln(err)
	}
}
