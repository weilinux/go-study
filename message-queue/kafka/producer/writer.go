package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"time"
)

// 上下文
var ctx = context.Background()

// 消息主题
var topic = "test_message001"
var topic2 = "test_message002"

// 创建时日志记录器
var l = logrus.New()

// 当前时间
var now = time.Now()

func main() {
	//TODO 使用segmentio/kafka-go操作
	// go get github.com/segmentio/kafka-go

	// kafka消息生产者配置
	writer := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"), // Kafka 服务地址列表
		Topic:        topic,                       // 生成消息的主题名称
		Balancer:     &kafka.LeastBytes{},         // 跨分区分发消息的平衡器(就按容其他第三方库或者语言)	默认使用循环分布
		MaxAttempts:  10,                          // 限制尝试发送消息的次数	默认最多尝试10次
		BatchSize:    100,                         // 限制在发送到分区之前缓冲多少消息	默认使用100条消息的目标批大小
		BatchBytes:   1048576,                     // 在发送到分区之前，以字节为单位限制请求的最大大小	默认值是使用卡夫卡默认值1048576。
		BatchTimeout: time.Second * 1,             // 刷新不完整消息批的频率的时间限制	默认值是至少每秒刷新一次
		ReadTimeout:  time.Second * 10,            // 读取操作超时	默认值为10秒
		WriteTimeout: time.Second * 10,            // 写入操作超时	默认值为10秒
		RequiredAcks: kafka.RequiredAcks(0),       // 在接收到对生产请求的响应之前，需要从分区副本确认的数量	RequireNone 0：不要等待 RequireOne 1：等待leader确认写入 RequireAll -1：等待完整的ISR确认写入	默认为RequireNone
		Async:        false,                       // 设置为true会导致写入消息方法不会被阻塞(错误被忽略，因为调用者将不会收到返回的值)	默认为false
		Completion: func(messages []kafka.Message, err error) {
			if err != nil {
				log.Fatalln(err)
			}
		}, // 当写入程序成功或失败时调用的可选函数
		Compression: kafka.Snappy,               // 设置用于压缩消息的压缩编解码器
		Logger:      kafka.LoggerFunc(l.Infof),  // 指定用于报告写入程序内部更改的记录器
		ErrorLogger: kafka.LoggerFunc(l.Errorf), // 用于报告错误的记录器
		Transport: &kafka.Transport{
			Dial: func(ctx context.Context, s string, s2 string) (net.Conn, error) {
				conn, err := kafka.Dial(s, s2)
				if err != nil {
					return nil, err
				}
				return conn, nil
			}, // 用于建立与卡夫卡群集的连接的函数
			DialTimeout: time.Second * 5,  // 设置与卡夫卡群集建立连接的时间限制	默认为5秒
			IdleTimeout: time.Second * 30, // 连接保持打开和未使用的最长时间	默认为30秒
			MetadataTTL: time.Second * 6,  // 此传输缓存的元数据的TTL 默认为6秒
			ClientID:    "client001",      // 传输在发送请求时与代理通信的唯一标识符
			TLS:         nil,              // 此传输建立的TLS连接的可选配置
			SASL:        nil,              // 将传输配置为使用SASL身份验证
			Resolver:    nil,              // 用于将代理主机名转换为网络地址的可选解析器
			Context:     ctx,              // 上下文
		}, // 用于向卡夫卡群集发送消息的传输 默认使用kafka.DefaultTransport
		AllowAutoTopicCreation: false, // 如果缺少主题则通知编写器创建主题
	}

	// 消息发送完毕后关闭writer
	defer func(w *kafka.Writer) {
		if err := w.Close(); err != nil {
			log.Fatalln(err)
		}
	}(writer)

	// 返回自上次调用方法以来，或自首次调用编写器时创建编写器以来，编写器统计信息的快照
	fmt.Println("详细信息：", writer.Stats())

	// 发送消息
	if err := writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,             // 消息主题
		Key:   []byte("key01"),   // 消息键名
		Value: []byte("value01"), // 消息值
		Headers: []kafka.Header{
			{
				Key:   "token",
				Value: []byte("dfg15sd1q5qwe15qw25"),
			},
		}, // 消息头部信息
		Time: now, // 消息时间
	}, kafka.Message{
		Topic:   topic2,
		Key:     []byte("key02"),
		Value:   []byte("value02"),
		Headers: nil,
		Time:    now,
	}); err != nil {
		log.Fatalln(err)
	}

}
