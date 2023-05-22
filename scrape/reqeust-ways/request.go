package main

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"log"
	"net/http"
)

// 指定采集的网址
var url = "http://www.baidu.com"

// 上下文
var ctx = colly.NewContext()

// 头部
var header = http.Header{
	"a": []string{"aaa", "bbb"},
	"c": []string{"ccc", "ddd"},
}

func main() {
	//TODO 使用gocolly/colly操作
	// go get github.com/gocolly/colly
	// 文档：http://go-colly.org

	// 创建采集器
	c := colly.NewCollector()

	// 初始化收集器的私有变量并设置收集器的默认配置
	c.Init()

	// get请求 如果 CheckHead=true 则在每次GET之前执行Head请求，以预先验证响应
	if err := c.Visit(url); err != nil {
		log.Fatalln(err)
	}

	// head请求
	if err := c.Head(url); err != nil {
		log.Fatalln(err)
	}

	// post请求
	requestData := make(map[string]string)
	requestData["a"] = "b"
	requestData["c"] = "e"
	if err := c.Post(url, requestData); err != nil {
		log.Fatalln(err)
	}

	// post（raw）请求
	requestJson, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalln(err)
	}
	if err = c.PostRaw(url, requestJson); err != nil {
		log.Fatalln(err)
	}

	// post（multipart/form-data）请求
	requestData2 := make(map[string][]byte)
	value1 := []byte{'a', 'b', 'c'}
	if jsonValue, err := json.Marshal(value1); err == nil {
		requestData2["a"] = jsonValue
	} else {
		log.Fatalln(err)
	}
	value2 := []byte{'d', 'e', 'f'}
	if jsonValue, err := json.Marshal(value2); err == nil {
		requestData2["a"] = jsonValue
	} else {
		log.Fatalln(err)
	}
	if err = c.PostMultipart(url, requestData2); err != nil {
		log.Fatalln(err)
	}

	// 自定义请求
	if err = c.Request("GET", url, nil, ctx, header); err != nil {
		log.Fatalln(err)
	}

	// 以序列化配置数据创建请求
	if request, err := c.UnmarshalRequest(nil); err == nil {
		// 提交请求
		if err = request.Do(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}
}
