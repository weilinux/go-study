package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	//1.创建路由
	r := gin.Default()

	//2.绑定路由规则和执行的函数  可设置为空作为入口导航页
	r.GET("", func(c *gin.Context) {
		//发送json格式数据
		c.JSON(http.StatusOK, gin.H{
			"message": "hello golang",
			"time":    time.Now().Format("2006-01-03 15:04:05"),
		})

		//发送string类型数据
		//c.String(http.StatusOK, "hello world")
	})

	//get请求 queryString
	r.GET("query", func(c *gin.Context) {
		//获取query参数 单参数 值：string 示例：?name=jack
		/*name, ok := c.GetQuery("name")
		if name == "" || !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "请输入名称",
				"name":    name,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"name":    name,
		})*/

		//获取query参数 单参数 值：数组 示例：?city=beijing&city=shanghai
		/*city := c.QueryArray("city")
		c.JSON(http.StatusOK, gin.H{"city": city})*/

		//获取query参数 多参数 值：string 示例：?info[name]=jack&info[age]=18&info[gender]=man
		info := c.QueryMap("info")
		c.JSON(http.StatusOK, gin.H{"info": info})
	})

	//get请求 路径变量
	r.GET("find/:name/:age", func(c *gin.Context) {
		//获取路径参数 单参数 值：string 示例：/jack/18
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
	})

	//post请求
	r.POST("create", func(c *gin.Context) {
		//获取post参数 单参数 值：string
		/*name, ok := c.GetPostForm("name")
		if name == "" || !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "请输入名称",
				"name":    name,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"name":    name,
		})*/

		//获取post参数 单参数 值：数组
		/*city := c.PostFormArray("city")
		c.JSON(http.StatusOK, gin.H{"city": city})*/

		//获取post参数 多参数 值：map
		info := c.PostFormMap("info")
		c.JSON(http.StatusOK, gin.H{"info": info})
	})

	//常规请求
	/*r.GET()
	r.POST()
	r.PUT()
	r.DELETE()
	r.PATCH()
	r.OPTIONS()
	r.HEAD()*/

	//多方式匹配，Get、Post、Put、Patch、Head、Options、Delete、Connect、Trace
	//r.Any()

	//未知路由处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found router"})
	})

	//未知调用方式
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found metchod"})
	})

	//3.监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
