package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 基于 cookie 的存储引擎，secret3652 参数是用于加密的密钥
//go get github.com/gin-contrib/sessions
var store = cookie.NewStore([]byte("secret"))

// 基于 Redis 存储 Session
//go get github.com/gin-contrib/sessions/redis
//store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

// 基于 MongoDB 存储 Session
//go get github.com/gin-contrib/sessions/mongo
//session, err := mgo.Dial("localhost:27017/test")
//c := session.DB("").C("sessions")
//store := mongo.NewStore(c, 3600, true, []byte("secret"))

// 基于 GoRM 存储 Session
//go get github.com/gin-contrib/sessions/gorm
//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//store := gorm.NewStore(db, true, []byte("secret"))

func main() {
	// TODO gin框架session
	// go get github.com/gin-contrib/sessions

	// 创建路由
	r := gin.Default()

	// 设置单个session
	r.Use(sessions.Sessions("mySession", store))

	// 设置多个session
	r.Use(sessions.SessionsMany([]string{"a", "b"}, store))

	//get请求
	r.GET("session", func(c *gin.Context) {
		//初始化 session 对象
		session := sessions.Default(c)

		// session_id
		//session.ID()

		// 删除session键
		//session.Delete(key)

		// 清空session
		//session.Clear()

		// session配置
		session.Options(sessions.Options{
			Path:     "/",         //cookie所在目录
			Domain:   "127.0.0.1", //cookie所在域 （作用范围）
			MaxAge:   3600,        //有效时长，单位秒
			Secure:   false,       //是否只能通过https访问
			HttpOnly: true,        //是否可以通过js代码进行操作
		})

		// 获取session的值
		if session.Get("key") != "value" {
			// 设置session的值
			session.Set("key", "value")

			// 保存session的值
			if err := session.Save(); err != nil {
				log.Fatalln(err)
			}
		}

		// 获取改定具体session名称的快捷方式（多个session模式）
		sessionA := sessions.DefaultMany(c, "a")
		if sessionA.Get("hello") != "world!" {
			sessionA.Set("hello", "world!")
			if err := sessionA.Save(); err != nil {
				log.Fatalln(err)
			}
		}

		sessionB := sessions.DefaultMany(c, "b")
		if sessionB.Get("hello") != "world?" {
			sessionB.Set("hello", "world?")
			if err := sessionA.Save(); err != nil {
				log.Fatalln(err)
			}
		}
		c.JSON(http.StatusOK, gin.H{"session的值是：": session.Get("key")})
	})

	// 监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
