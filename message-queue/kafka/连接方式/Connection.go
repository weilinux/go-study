package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

// 上下文
var ctx = context.Background()

// 网络协议
var network = "tcp"

// 地址
var address = "localhost:9092"

// 消息主题
var topic = "test_message001"

// 消息分区
var partition = 0

func main() {
	//TODO 使用segmentio/kafka-go操作
	// go get github.com/segmentio/kafka-go

	// 连接方式
	if conn, err := kafka.Dial(network, address); err == nil {
		if err = conn.Close(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

	// 指定上下文的连接
	if conn, err := kafka.DialContext(ctx, network, address); err == nil {
		if err = conn.Close(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

	// 指定上下文、消息主题、消息分区的连接
	if conn, err := kafka.DialLeader(ctx, network, address, topic, partition); err == nil {
		if err = conn.Close(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

	// 定义区分的连接
	conn, err := kafka.DialPartition(ctx, network, address, kafka.Partition{
		Topic: topic, // 消息主题
		ID:    0,     // 索引
		Leader: kafka.Broker{
			Host: "localhost",    // 域名
			Port: 9092,           // 端口
			ID:   0,              // 代理id
			Rack: "rack_test_01", // 指定broker机架信息 若设置了机架信息，kafka在分配副本时会考虑把某个分区的多个副本分配在多个机架上，这样即使某个机架上的broker全部崩溃，也能保证其他机架上的副本可以正常工作
		}, // kafka集群中的kafka代理
		Replicas: nil,
		Isr:      nil,
		Error:    nil,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// 关闭连接
	defer func(conn *kafka.Conn) {
		if err = conn.Close(); err != nil {
			log.Fatalln(err)
		}
	}(conn)

	// 指定连接、消息主题、消息分区创建出的新连接(NewConnWith简化版)
	conn2 := kafka.NewConn(conn, topic, partition)
	defer func(conn2 *kafka.Conn) {
		if err = conn2.Close(); err != nil {
			log.Fatalln(err)
		}
	}(conn2)

	// 指定连接、配置对象创建出的新连接(完整版)
	conn3 := kafka.NewConnWith(conn, kafka.ConnConfig{
		ClientID:        "client001",    // 传输在发送请求时与代理通信的唯一标识符
		Topic:           topic,          // 消息主题
		Partition:       partition,      // 消息分区
		Broker:          0,              // 代理id
		Rack:            "rack_test_01", // broker机架信息
		TransactionalID: "t_id_01",      // 用于事务传递的事务id。如果配置了事务id，则应启用幂等传递
	})
	defer func(conn3 *kafka.Conn) {
		if err = conn3.Close(); err != nil {
			log.Fatalln(err)
		}
	}(conn3)
}
