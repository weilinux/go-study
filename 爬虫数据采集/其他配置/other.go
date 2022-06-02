package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/proxy"
	"github.com/gocolly/colly/storage"
	"golang.org/x/net/publicsuffix"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	//TODO 使用gocolly/colly操作
	// go get github.com/gocolly/colly
	// 文档：http://go-colly.org

	// 创建采集器
	c := colly.NewCollector()

	// 初始化收集器的私有变量并设置收集器的默认配置
	c.Init()

	// 请求重定向管理
	c.RedirectHandler = func(req *http.Request, via []*http.Request) error {
		// handler body...
		return nil
	}

	// 在每次GET之前执行Head请求，以预先验证响应
	c.CheckHead = true

	// 克隆创建收集器的精确副本
	c2 := c.Clone()
	c2.ID = 002

	// 将调试器附加到收集器
	c.SetDebugger(&debug.LogDebugger{})

	// 设置自定义http配置
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment, // 返回用于给定请求的代理的URL
		DialContext: net.Dialer{ //	拨号器包含连接到地址的选项
			Timeout:   time.Second * 30, // 拨号等待连接完成的最长时间（超时）
			KeepAlive: time.Second * 30, // 指定活动网络连接的保持活动探测之间的间隔
		}.DialContext,
		DialTLSContext:         nil,   // 可选的拨号函数
		TLSClientConfig:        nil,   // 用于tls.Client的TLS配置
		TLSHandshakeTimeout:    0,     // 等待tls握手的最长时间
		DisableKeepAlives:      false, // 如果为true，则禁用HTTP keep-alives，并将只对单个HTTP请求使用到服务器的连接
		DisableCompression:     false, // 如果为true，则当请求不包含现有的Accept-Encoding值时，禁止传输使用“Accept-Encoding: gzip”请求头请求压缩。如果传输本身请求gzip并得到gzip响应，它会在Response.Body中被透明解码
		MaxIdleConns:           100,   // 最大空闲连接数
		MaxIdleConnsPerHost:    20,    // 最大idl(保持活动)连接数
		MaxConnsPerHost:        0,     // 限制连接总数，包括处于拨号、活动和空闲状态的连接 违反限制时，刻度盘将被阻止
		IdleConnTimeout:        0,     // 空闲(保持活动)连接在关闭之前保持空闲的最长时间
		ResponseHeaderTimeout:  0,     // 指定在完全写入请求(包括请求体，如果有)后等待服务器响应头的时间。这个时间不包括阅读响应正文的时间
		ExpectContinueTimeout:  0,     // 在请求具有“Expect: 100-continue”标头的情况下，完全写入请求标头后等待服务器的第一个响应标头的时间
		TLSNextProto:           nil,   // 指定在tls ALPN协议协商后传输如何切换到备用协议(如HTTP/2)
		ProxyConnectHeader:     nil,   // 指定在连接请求期间发送给代理的头。若要动态设置标头
		GetProxyConnectHeader:  nil,   // 指定一个func来返回在对ip:port目标的连接请求期间发送给proxyURL的头
		MaxResponseHeaderBytes: 0,     // 指定服务器响应头中允许的响应字节数的限制
		WriteBufferSize:        0,     // 指定写入传输时使用的写入缓冲区的大小
		ReadBufferSize:         0,     // 指定从传输中读取时使用的读取缓冲区的大小
		ForceAttemptHTTP2:      false, // 控制在非零值时是否启用HTTP/2
	})

	// 关闭cookie处理
	c.DisableCookies()

	// 根据域名安全地设置cookies
	if jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List}); err == nil {
		// 设置cookiejar 应用程序（浏览器或非浏览器）在其中放置它在请求和响应期间使用的 cookie
		c.SetCookieJar(jar)
	} else {
		log.Fatalln(err)
	}

	// 返回指定URL的请求中要发送的cookie
	fmt.Println("要发送的cookie：", c.Cookies("xxx.com"))

	// 设置时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalln(err)
	}
	// 将字符串转成time
	expiresTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-06-30 23:59:59", location)
	if err != nil {
		log.Fatalln(err)
	}
	// 设置登录cookies
	if err = c.SetCookies("xxx.com", []*http.Cookie{
		{
			Name:     "remember_user_token",                                      // cookie名称
			Value:    "wNDUxOV0sIiQyYSQxMSRwdkhqWVhHYmxXaDJ6dEU3NzJwbmsuIiwiMTU", // cookie值
			Path:     "/",                                                        // 可以访问该cookie的路径
			Domain:   ".xxx.com",                                                 // 可以访问该该cookie的域名
			Secure:   true,                                                       // Cookie只能通过被HTTPS协议加密过的请求发送给服务端
			Expires:  expiresTime,                                                // 过期时间，绝对时间（某一个具体时间）
			HttpOnly: true,                                                       // 主要是禁止调用js的document.cookie这个API从而防止跨站脚本攻击（XSS）
		},
	}); err != nil {
		log.Fatalln(err)
	}

	// 设置收集器的默认超时 （默认10s）
	c.SetRequestTimeout(time.Second * 30)

	// 创建默认内存存储器
	inMemoryStorage := &storage.InMemoryStorage{}
	// 初始化存储器
	if err := inMemoryStorage.Init(); err == nil {
		// 设置采集器的默认存储器
		if err = c.SetStorage(inMemoryStorage); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

	// 为收集器设置代理 代理类型由URL方案确定 支持“http”和“socks5”（如果方案为空则默认为http）
	if err := c.SetProxy("socks5://127.0.0.1:1234"); err != nil {
		log.Fatalln(err)
	}

	// 内置的通过轮询方式实现代理切换的函数
	if switcher, err := proxy.RoundRobinProxySwitcher(
		"socks5://127.0.0.1:1235",
		"http://127.0.0.1:1237",
	); err == nil {
		// 设置自定义代理切换器函数
		c.SetProxyFunc(switcher)
	} else {
		log.Fatalln(err)
	}

	// 配置域的链接限制（限速）
	limitRule := &colly.LimitRule{
		DomainRegexp: "",              // 与域匹配的正则表达式
		DomainGlob:   "*.www.com/*",   // 与域匹配的全局模式
		Delay:        3 * time.Second, // 向匹配域创建新请求之前等待的时间
		RandomDelay:  1 * time.Second, // 创建新请求之前添加到延迟中的额外随机等待时间
		Parallelism:  10,              // 匹配域允许的最大并发请求数
	}
	// 初始化配置
	if err := limitRule.Init(); err == nil {
		// 为采集器设置限速配置
		if err = c.Limit(limitRule); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

	// 创建多个限速配置
	limitRules := make([]*colly.LimitRule, 3, 3)
	// 配置一
	limitRules[0] = &colly.LimitRule{
		DomainRegexp: "",              // 与域匹配的正则表达式
		DomainGlob:   "*.www1.com/*",  // 与域匹配的全局模式
		Delay:        3 * time.Second, // 向匹配域创建新请求之前等待的时间
		RandomDelay:  1 * time.Second, // 创建新请求之前添加到延迟中的额外随机等待时间
		Parallelism:  10,              // 匹配域允许的最大并发请求数
	}
	if err := limitRules[0].Init(); err != nil {
		log.Fatalln(err)
	}
	// 配置二
	limitRules[1] = &colly.LimitRule{
		DomainRegexp: "",              // 与域匹配的正则表达式
		DomainGlob:   "*.www2.com/*",  // 与域匹配的全局模式
		Delay:        3 * time.Second, // 向匹配域创建新请求之前等待的时间
		RandomDelay:  1 * time.Second, // 创建新请求之前添加到延迟中的额外随机等待时间
		Parallelism:  10,              // 匹配域允许的最大并发请求数
	}
	if err := limitRules[1].Init(); err != nil {
		log.Fatalln(err)
	}
	// 配置三
	limitRules[2] = &colly.LimitRule{
		DomainRegexp: "",              // 与域匹配的正则表达式
		DomainGlob:   "*.www3.com/*",  // 与域匹配的全局模式
		Delay:        3 * time.Second, // 向匹配域创建新请求之前等待的时间
		RandomDelay:  1 * time.Second, // 创建新请求之前添加到延迟中的额外随机等待时间
		Parallelism:  10,              // 匹配域允许的最大并发请求数
	}
	if err := limitRules[2].Init(); err != nil {
		log.Fatalln(err)
	}
	// 为采集器设置多个限速配置
	if err := c.Limits(limitRules); err != nil {
		log.Fatalln(err)
	}

	// 包含有关收集器内部的有用调试信息
	fmt.Println("收集器的文本表示形式：", c.String())
}
