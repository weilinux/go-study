package main

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//上下文
var ctx = context.Background()

func main() {
	//TODO 使用gorm操作
	//go get -u gorm.io/gorm
	//go get gorm.io/driver/postgres

	//连接mysql
	db, err := gorm.Open(postgres.New(postgres.Config{ //自定义驱动
		DriverName: "cloudsqlpostgres",                                                                             // 自定义驱动，通过 DriverName 选项自定义 MySQL 驱动
		DSN:        "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable", // DSN data source name
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
