package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Cipher   string `json:"cipher"`
}

type Param struct {
	Name string `form:"name" json:"name"`
	Age  string `form:"age" json:"age"`
	Uid  string `json:"uid"`
}

func (u *User) encrypt() {
	h := md5.New()
	h.Write([]byte(u.Password))
	u.Cipher = hex.EncodeToString(h.Sum(nil))
}

func (p *Param) findUser() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	p.Uid = string(bytes)
}

func main() {
	//创建路由
	r := gin.Default()

	//Shouldxxx和bindxxx区别就是bindxxx会在head中添加400的返回信息，而Shouldxxx不会 (400: bad request)

	//绑定header
	//c.ShouldBindHeader()

	//绑定url
	//c.ShouldBindUri()

	//绑定查询
	//c.ShouldBindQuery()

	//绑定表单
	//ShouldBind

	//绑定json
	//c.ShouldBindJSON()

	//绑定XML
	//c.ShouldBindXML()

	//get请求 路径变量
	r.GET("findUser/:name/:age", func(c *gin.Context) {
		//申明结构体
		var param Param

		//参数绑定 query
		if err := c.ShouldBindQuery(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//查询uid
		param.findUser()

		//返回结构体json格式数据
		c.JSON(http.StatusOK, gin.H{"result": param.Uid})
	})

	//post请求
	r.POST("addUser", func(c *gin.Context) {
		//申明结构体
		var user User

		//参数绑定 form
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//加密
		user.encrypt()

		//返回结构体json格式数据
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
