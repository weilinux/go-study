package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func logMiddleware() gin.HandlerFunc {
	// 创建log实例
	var logger = logrus.New()

	// 以JSON而不是默认的ASCII格式记录
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05", // 定义时间格式
		DisableTimestamp:  false,                 // 输出中禁止自动写入时间戳
		DisableHTMLEscape: false,                 // 输出中禁用html转义
		DataKey:           "data",                // 将自定义字段放在给定的dataKey中
		PrettyPrint:       true,                  // 缩进所有json日志
	})

	// 记录所有级别的错误
	logger.SetLevel(logrus.TraceLevel)

	// 日志定位行号 如：func=main.main file="./xxx.go:38"
	logger.SetReportCaller(true)

	// 已追加模式打开时可以更安全的写入，自动调用互斥锁（线程安全）
	logger.SetNoLock()

	// 获取当前时间
	now := time.Now()

	// 目录地址
	dirPath := path.Join("log", "api_log", strconv.Itoa(now.Year()), strconv.Itoa(int(now.Month())))

	// 如果目录不存在则逐层创建
	_, err := os.Stat(dirPath)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0644)
		if err != nil {
			logger.Errorln("create dir err info:", err)
		}
	}

	// 文件地址
	filePath := path.Join(dirPath, fmt.Sprintf("%s.log", strconv.Itoa(now.Day())))

	// 打开文件
	src, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.Errorln("create file err info:", err)
	}

	// 将日志输出在文件中
	logger.SetOutput(src)

	// 写入请求日志
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 状态码
		statusCode := c.Writer.Status()
		// 请求ip
		clientIP := c.ClientIP()
		// 真实访问IP（nginx代理）
		realIP := c.GetHeader("X-Real-IP")
		// 获取授权认证信息
		auth := c.GetHeader("Authorization")
		// 获取query参数
		queryData, _ := c.Get("query_data")
		// 获取form参数
		formData, _ := c.Get("form_data")

		// 日志格式
		logger.WithFields(logrus.Fields{
			"start_time":   startTime.Format("2006-01-02 15:04:05"),
			"end_time":     endTime.Format("2006-01-02 15:04:05"),
			"latency_time": latencyTime,
			"req_method":   reqMethod,
			"status_code":  statusCode,
			"client_ip":    clientIP,
			"real_ip":      realIP,
			"auth":         auth,
			"query_data":   queryData,
			"form_data":    formData,
		}).Info("http request info")
	}
}

func main() {
	//TODO 通过gin框架创建logrus中间件访问日志

	//创建路由
	r := gin.Default()

	//注册全局中间件
	r.Use(logMiddleware())

	//get请求 路径变量
	r.GET("/find/:name/:age", func(c *gin.Context) {
		//创建变量
		queryData := gin.H{}

		//获取路径参数 单参数 值：string 示例：/jack/18
		queryData["name"] = c.Param("name")
		queryData["age"] = c.Param("age")

		//将请求参数写入到上下文中
		c.Set("query_data", queryData)

		//返回结果
		c.JSON(http.StatusOK, queryData)
	})

	//post请求
	r.POST("create", func(c *gin.Context) {
		//创建变量
		formData := gin.H{}

		//获取post参数 单参数 值：string
		var ok bool
		formData["name"], ok = c.GetPostForm("name")
		if formData["name"] == "" || !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "请输入名称",
				"name":    formData["name"],
			})
			return
		}

		//获取post参数 单参数 值：数组
		formData["city"] = c.PostFormArray("city")

		//获取post参数 多参数 值：map
		formData["info"] = c.PostFormMap("info")

		//将请求参数写入到上下文中
		c.Set("form_data", formData)

		//返回结果
		c.JSON(http.StatusOK, formData)
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
