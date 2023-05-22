package main

import (
	"context"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//上下文
var ctx = context.Background()

func main() {
	//TODO 使用gorm操作
	//go get -u gorm.io/gorm
	//go get gorm.io/driver/mysql

	//连接mysql
	db, err := gorm.Open(sqlserver.New(sqlserver.Config{ //自定义驱动
		DriverName:        "my_mysql_driver",                                                          // 自定义驱动，通过 DriverName 选项自定义 MySQL 驱动
		DSN:               "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 256,                                                                        // string 类型字段的默认长度
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//持续会话模式通常被用于执行一系列操作
	tx := db.WithContext(ctx)

	//CURD...
	tx.Create("")
	tx.Find("")
	tx.Updates("")
	tx.Delete("")

	//详细操作观看：https://gorm.io/zh_CN/docs/v2_release_note.html
}
