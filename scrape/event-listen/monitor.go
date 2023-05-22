package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"log"
)

// 指定采集的网址
var url = "http://www.baidu.com"

func main() {
	//TODO 使用gocolly/colly操作
	// go get github.com/gocolly/colly
	// 文档：http://go-colly.org

	/**
	回掉函数的调用顺序如下：
		1、OnRequest 在发起请求前被调用
		2、OnError 请求过程中如果发生错误被调用
		3、OnResponse 收到回复后被调用
		4、OnHTML 在OnResponse之后被调用，如果收到的内容是HTML
		5、OnScraped 在OnHTML之后被调用
	*/

	// 创建采集器
	c := colly.NewCollector()

	// 初始化收集器的私有变量并设置收集器的默认配置
	c.Init()

	// 请求执行之前调用
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	// 响应返回之后调用
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Visited", response.Request.URL)
	})

	// 配置节点信息
	goquerySelector := &goquery.Selection{
		Nodes: []*html.Node{
			0: {
				Type:      html.DocumentNode,
				DataAtom:  0,
				Data:      "",
				Namespace: "",
				Attr:      nil,
			},
		},
	}
	// 采集结构体信息
	type Nested struct {
		String  string   `selector:"div > p"`
		Classes []string `selector:"li" attr:"class"`
		Struct  *Nested  `selector:"div > div"`
	}
	nested := &Nested{
		String:  "",
		Classes: nil,
		Struct:  nil,
	}
	// 采集信息 自定义
	if err := colly.UnmarshalHTML(nested, goquerySelector); err != nil {
		log.Fatalln(err)
	}

	// 监听执行 selector（函数将在GoQuery选择器参数匹配的每个HTML元素上执行） html选择器类似于jQuery
	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		if err := element.Request.Visit(element.Attr("href")); err != nil {
			log.Fatalln(err)
		}
	})
	c.OnHTML("tr td:nth-of-type(1)", func(element *colly.HTMLElement) {
		fmt.Println("First column of a table row:", element.Text)
	})

	// 监听执行 selector（函数将在xpath查询参数匹配的每个XMLelement上执行）
	//c.OnXML()

	// 取消监听 HTML元素
	c.OnHTMLDetach("input")

	// 取消监听 XMLelement
	//c.OnXMLDetach()

	// 错误回调
	c.OnError(func(response *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// 在OnHMTL之后执行，作为抓取的最后一步
	c.OnScraped(func(response *colly.Response) {
		fmt.Println("Finished", response.Request.URL)
	})

	// get请求 如果 CheckHead=true 则在每次GET之前执行Head请求，以预先验证响应
	if err := c.Visit(url); err != nil {
		log.Fatalln(err)
	}
}
