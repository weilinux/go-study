package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/storage"
	"log"
	"net/http"
	"net/url"
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

	// 创建默认内存存储器
	inMemoryStorage := &storage.InMemoryStorage{}

	// 设置cookies
	cookies := []*http.Cookie{
		{
			Name:     "remember_user_token",                                      // cookie名称
			Value:    "wNDUxOV0sIiQyYSQxMSRwdkhqWVhHYmxXaDJ6dEU3NzJwbmsuIiwiMTU", // cookie值
			Path:     "/",                                                        // 可以访问该cookie的路径
			Domain:   ".xxx.com",                                                 // 可以访问该该cookie的域名
			Secure:   true,                                                       // Cookie只能通过被HTTPS协议加密过的请求发送给服务端
			Expires:  time.Now(),                                                 // 过期时间，绝对时间（某一个具体时间）
			HttpOnly: true,                                                       // 主要是禁止调用js的document.cookie这个API从而防止跨站脚本攻击（XSS）
		},
	}
	// 序列化 http.cookie 列表
	stringifyCookies := storage.StringifyCookies(cookies)

	// 反序列化 cookies 字符串
	storage.UnstringifyCookies(stringifyCookies)

	// 检查cookie中是否表示cookie名称
	fmt.Println("改cookie是否存在于cookis中：", storage.ContainsCookie(cookies, "_csrf"))

	// 初始化存储器
	if err := inMemoryStorage.Init(); err != nil {
		log.Fatalln(err)
	}

	// 关闭存储器
	defer func() {
		if err := inMemoryStorage.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// 访问该标识事件的HTTP请求
	if err := inMemoryStorage.Visited(1); err != nil {
		log.Fatalln(err)
	}

	// 是否已访问该标识事件的HTTP请求
	if visited, err := inMemoryStorage.IsVisited(2); err == nil {
		fmt.Println("是否访问：", visited)
	} else {
		log.Fatalln()
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
	inMemoryStorage.SetCookies(u, "_ga=GA1.2.1611472128.1650815524; _gid=GA1.2.2080811677.1652022429; __atuvc=2|17,0|18,5|19")

	// 获取存储器对应url的cookie
	fmt.Println("cookie信息：", inMemoryStorage.Cookies(u))

	// 设置采集器的默认存储器
	if err := c.SetStorage(inMemoryStorage); err != nil {
		log.Fatalln(err)
	}
}
