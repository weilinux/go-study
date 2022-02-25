package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//基于 cookie 的存储引擎，secret3652 参数是用于加密的密钥
//go get github.com/gin-contrib/sessions
var store = cookie.NewStore([]byte("secret3652"))

//基于 Redis 存储 Session
//go get github.com/gin-contrib/sessions/redis
//store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret3652"))

func main() {
	//创建路由
	r := gin.Default()

	//设置 session 中间件，参数 mySession，指的是 session 的名字，也是 cookie 的名字
	//store 是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mySession", store))

	//get请求
	r.GET("session", func(c *gin.Context) {
		//初始化 session 对象
		session := sessions.Default(c)

		//session_id
		//session.ID()

		//删除session键
		//session.Delete(key)

		//清空session
		//session.Clear()

		//session配置
		session.Options(sessions.Options{
			Path:     "/",         //cookie所在目录
			Domain:   "127.0.0.1", //cookie所在域 （作用范围）
			MaxAge:   3600,        //有效时长，单位秒
			Secure:   false,       //是否只能通过https访问
			HttpOnly: true,        //是否可以通过js代码进行操作
		})

		//获取session的值
		if session.Get("key") != "value" {
			//设置session的值
			session.Set("key", "value")

			//保存session的值
			if err := session.Save(); err != nil {
				log.Fatalln(err)
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"session的值是：": session.Get("key")})
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
