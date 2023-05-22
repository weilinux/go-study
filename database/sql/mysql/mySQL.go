package mysql

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//上下文
var ctx = context.Background()

func main() {
	//TODO 使用gorm操作
	//go get -u gorm.io/gorm
	//go get gorm.io/driver/mysql

	//连接mysql
	db, err := gorm.Open(mysql.New(mysql.Config{ //自定义驱动
		DriverName:                "my_mysql_driver",                                                          // 自定义驱动，通过 DriverName 选项自定义 MySQL 驱动
		DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//持续会话模式通常被用于执行一系列操作
	tx := db.WithContext(ctx)

	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, err := tx.DB()

	// Ping
	sqlDB.Ping()

	// Close
	defer sqlDB.Close()

	// 返回数据库统计信息
	sqlDB.Stats()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//CURD...
	tx.Create("")
	tx.Find("")
	tx.Updates("")
	tx.Delete("")

	//详细操作观看：https://gorm.io/zh_CN/docs/v2_release_note.html
}
