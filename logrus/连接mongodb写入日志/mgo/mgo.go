package main

import (
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
)

// 创建log实例
var logger = logrus.New()

func main() {
	//TODO 使用sirupsen/logrus自定义hook
	// go get github.com/weekface/mgorus

	// 无用户名、密码
	//hooker, err := mgorus.NewHooker("127.0.0.1:27017", "db", "collection")

	// 需要用户名、密码
	//hooker, err := mgorus.NewHookerWithAuth("127.0.0.01:27017", "db", "collection", "user", "pass")

	// 副本集需要权限表、用户名、密码
	//hooker, err := mgorus.NewHookerWithAuthDb("127.0.0.01:27017", "authdb", "db", "collection", "user", "pass")

	// 使用内置mongo连接
	hooker, err := mgorus.NewHooker("192.168.17.128:27017", "cbec_rbac", "test")
	if err != nil {
		logger.Error(err)
	} else {
		logger.Hooks.Add(hooker)
	}

	// 使用mgo连接
	/*s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"localhost:27017"}, //地址
		Timeout:  5 * time.Second,             //过期时间
		Database: "admin",                     //数据库
		Username: "",                          //用户名
		Password: "",                          //密码
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), &tls.Config{InsecureSkipVerify: true})
			return conn, err
		},
	})
	if err != nil {
		logger.Error("can't create session: %s\n", err)
	}

	// 指定db和集合
	c := s.DB("db").C("collection")

	// 生成钩子
	hooker := mgorus.NewHookerFromCollection(c)*/

	// 将钩子添加到logger实例
	logger.Hooks.Add(hooker)

	entry := logger.WithFields(logrus.Fields{
		"name":  "jack",
		"age":   18,
		"email": "123@qq.com",
	})
	entry.Warning("warning write error")
	entry.Fatal("fatal write error")
}
