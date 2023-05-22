package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"time"
)

// 上下文
var ctx = context.Background()

// 消息主题
var topic = "test_message001"

// 创建时日志记录器
var l = logrus.New()

func main() {
	//TODO 使用segmentio/kafka-go操作
	// go get github.com/segmentio/kafka-go

	// kafka消息消费者配置
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:                []string{"localhost:9092"}, // 用于连接kafka群集的代理地址列表
		GroupID:                "",                         // 可选的使用者组id	如果指定了GroupID，则不应指定分区
		GroupTopics:            nil,                        // 指定多个主题，但只能与GroupID结合使用
		Topic:                  topic,                      // 要从中读取消息的主题
		Partition:              0,                          // 要从中读取消息的分区	只能二选一分配分区或组ID
		Dialer:                 nil,                        // 打开与kafka服务器的连接的拨号程序
		QueueCapacity:          100,                        // 内部消息队列的容量 默认值为100
		MinBytes:               1,                          // 向代理指示使用者将接受的最小批大小	默认值为1
		MaxBytes:               1024 * 1024 * 1000,         // 向代理指示使用者将接受的最大批大小	默认值为1MB
		MaxWait:                time.Second * 10,           // 等待新数据到来的最长时间 默认为10秒
		ReadLagInterval:        time.Second * 3,            // 设置读取器延迟的更新频率 将此字段设置为负值将禁用滞后报告
		GroupBalancers:         nil,                        // 将提供给协调器的客户端消费者组平衡策略的优先级排序列表
		HeartbeatInterval:      time.Second * 3,            // 设置读取器发送消费者组心跳更新的可选频率 默认为3秒
		CommitInterval:         time.Second * 0,            // 将偏移提交给代理的间隔 如果为0，则将同步处理提交 默认值为0
		PartitionWatchInterval: time.Second * 5,            // 读取器检查分区更改的频率 如果读取器看到分区更改（如分区添加），它将重新平衡选择新分区的组 默认为5秒
		WatchPartitionChanges:  false,                      // 主题发生任何分区更改，消费者团体应该轮询代理并重新平衡
		SessionTimeout:         time.Second * 30,           // 选择设置协调器认为消费者已死亡并启动重新平衡之前可能经过的没有心跳的时间长度 默认为30秒
		RebalanceTimeout:       time.Second * 30,           // 选择设置协调器等待成员加入的时间长度，作为重新平衡的一部分 默认为30秒
		JoinGroupBackoff:       time.Second * 5,            // 选择设置出错后重新加入使用者组之间的等待时间长度 默认为5秒
		RetentionTime:          time.Hour * 24,             // 选择设置代理保存消费者组的时间长度 默认为24小时
		StartOffset:            kafka.FirstOffset,          // 当消费者组找到没有提交偏移量的分区时，应该从何处开始消费 默认值FirstOffset
		ReadBackoffMin:         time.Millisecond * 100,     // 选择设置读取器在轮询新消息之前等待的最小时间 默认为100毫秒
		ReadBackoffMax:         time.Second * 1,            // 选择设置读取器在轮询新消息之前等待的最长时间 默认为1秒
		Logger:                 kafka.LoggerFunc(l.Infof),  // 指定用于报告读取器内部更改的记录器
		ErrorLogger:            kafka.LoggerFunc(l.Errorf), // 用于报告错误的记录器
		IsolationLevel:         kafka.ReadUncommitted,      // 控制事务记录的可见性	ReadUncommitted使所有记录可见；对于ReadCommitted，只有非事务记录和已提交记录可见
		MaxAttempts:            3,                          // 在传递错误之前进行的尝试次数限制 默认值为3次
		OffsetOutOfRangeError:  false,                      // 表示如果发生OffsetAutoFrange错误，读取器应该返回错误，而不是无限期地重试
	})

	// 消息读取完毕后关闭reader
	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(reader)

	// 返回自上次调用该方法以来，或自首次调用读取器时创建读取器以来，读取器Stats的快照
	fmt.Println("详细信息：", reader.Stats())

	// 读取器配置信息
	fmt.Println("配置信息：", reader.Config())

	// 返回ReadMessage返回的最后一条消息的延迟，如果r由使用者组支持，则返回-1
	fmt.Println("延迟为：'", reader.Lag())

	// 通过获取主题和分区的最后一个偏移量并计算该值与ReadMessage返回的最后一条消息的偏移量之间的差值，返回读取器的当前延迟
	if lag, err := reader.ReadLag(ctx); err == nil {
		fmt.Println("当前延迟为：'", lag)
	} else {
		log.Fatalln(err)
	}

	// 返回读取器的当前绝对偏移量，如果r由使用者组支持，则返回-1
	fmt.Println("当前绝对偏移量：'", reader.Offset())

	// 更改将从中读取下一批消息的偏移量。该方法因io而失败。如果读取器已关闭，则返回ErrClosedPipe
	if err := reader.SetOffset(kafka.FirstOffset); err == io.ErrClosedPipe {
		fmt.Println("读取器已关闭")
	} else {
		log.Fatalln(err)
	}

	// 更改给定时间戳t时读取下一批消息的偏移量
	if err := reader.SetOffsetAt(ctx, time.Now().Add(time.Second*60)); err != nil {
		log.Fatalln(err)
	}

	// 读取并返回来自r的下一条消息。该方法调用会一直阻塞，直到消息可用或发生错误为止。程序还可以指定上下文以异步取消阻塞操作
	// 使用使用者组时，FetchMessage不会自动提交偏移量
	// 使用CommitMessages提交偏移量
	if message, err := reader.FetchMessage(ctx); err == nil {
		fmt.Println(message)

		// 提交作为参数传递的消息列表。当程序配置为阻塞时，程序可以传递上下文以异步取消提交操作
		if err = reader.CommitMessages(ctx, message); err != nil {
			log.Fatalln(err)
		}
	} else {
		if err == io.EOF {
			fmt.Println("读取器已关闭")
		} else {
			log.Fatalln(err)
		}
	}

	// 读取并返回下一条消息。该方法调用会一直阻塞，直到消息可用或发生错误 程序还可以指定上下文以异步取消阻塞操作
	// 使用死循环读取并打印每条信息
	for {
		if message, err := reader.ReadMessage(ctx); err == nil {
			fmt.Println(message)
		} else {
			if err == io.EOF {
				fmt.Println("读取器已关闭")
			} else {
				log.Fatalln(err)
			}
		}
	}
}
