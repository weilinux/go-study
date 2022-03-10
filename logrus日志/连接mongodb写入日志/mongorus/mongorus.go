package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wo4zhuzi/mongorus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 创建log实例
var logger = logrus.New()

func main() {
	//TODO 使用wo4zhuzi/mongorus自定义hook
	// go get github.com/wo4zhuzi/mongorus

	// 使用mongo-driver连接
	hooker, err := mongorus.NewAuthMongoHook("127.0.0.1:12017", "test_db", "test_collection", options.Credential{
		Username: "test_username",
		Password: "test_password",
	})
	if err == nil {
		// 将钩子添加到logger实例
		logger.Hooks.Add(hooker)
	} else {
		fmt.Print(err)
	}

	entry := logger.WithFields(logrus.Fields{
		"name":  "jack",
		"age":   18,
		"email": "123@qq.com",
	})
	entry.Warning("warning write error")
	entry.Fatal("fatal write error")
}
