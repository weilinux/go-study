package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"io"
	"log"
	"time"
)

// 上下文
var ctx = context.Background()

// 网络协议
var network = "tcp"

// 地址
var address = "localhost:9092"

// 消息主题
var topic = "test_message001"
var topic2 = "test_message002"

// 读取信息
var rb = make([]byte, 1024)

// 写入信息
var wb = make([]byte, 1024)

// 当前时间
var now = time.Now()

// 写入信息
var m = kafka.Message{
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
}
var m2 = kafka.Message{
	Topic:   topic2,
	Key:     []byte("key02"),
	Value:   []byte("value02"),
	Headers: nil,
	Time:    now,
}

func main() {
	//TODO 使用segmentio/kafka-go操作
	// go get github.com/segmentio/kafka-go

	// 连接方式
	conn, err := kafka.Dial(network, address)
	if err != nil {
		log.Fatalln(err)
	}

	// 关闭连接
	defer func(conn *kafka.Conn) {
		if err = conn.Close(); err != nil {
			log.Fatalln(err)
		}
	}(conn)

	// 返回当前建立连接的代理信息
	fmt.Println("代理信息", conn.Broker())

	// 返回控制器当前请求的代理信息
	if broker, err := conn.Controller(); err == nil {
		fmt.Println("代理信息", broker)
	} else {
		log.Fatalln(err)
	}

	// 返回所有代理信息
	if brokers, err := conn.Brokers(); err == nil {
		for _, broker := range brokers {
			fmt.Println("代理信息", broker)
		}
	} else {
		log.Fatalln(err)
	}

	// 删除指定的主题
	if err = conn.DeleteTopics(topic, topic2); err != nil {
		log.Fatalln(err)
	}

	// 返回本地网络地址
	fmt.Println("本地网络地址：", conn.LocalAddr())

	// 返回远程网络地址
	fmt.Println("远程网络地址：", conn.RemoteAddr())

	// 设置与连接关联的读写截止日期
	if err = conn.SetDeadline(now.Add(time.Hour * 1)); err != nil {
		log.Fatalln(err)
	}

	// 设置未来读取调用和任何当前阻止的读取调用的截止日期	如果 t 等于零值表示读取不会超时
	if err = conn.SetReadDeadline(now.Add(time.Hour * 1)); err != nil {
		log.Fatalln(err)
	}

	// 设置未来写调用和任何当前阻止的写调用的截止日期(即使写入超时，它也可能返回n>0，表示部分数据已成功写入)	如果 t 等于零值表示写入不会超时
	if err = conn.SetWriteDeadline(now.Add(time.Hour * 1)); err != nil {
		log.Fatalln(err)
	}

	// 以整数对的形式返回连接的当前偏移量，其中第一个是偏移量值，第二个指示相对于偏移量进行搜索
	offset, whence := conn.Offset()
	fmt.Println("偏移量：", offset, "相对于偏移量进行搜索：", whence)

	// 设置下一次读写操作的偏移量	返回连接的新绝对偏移量
	if offset, err = conn.Seek(kafka.FirstOffset, kafka.SeekStart); err == nil {
		fmt.Println("偏移量：", offset)
	} else {
		log.Fatalln(err)
	}

	// 以连接的当前偏移量读取消息，成功时将偏移量提前，以便对Read方法下一次调用将生成下一条消息
	if n, err := conn.Read(rb); err == nil {
		fmt.Println("读取内容：", string(rb), "读取字节数：", n)
	} else {
		if err == io.ErrShortBuffer {
			fmt.Println("缓冲区已写满")
		} else {
			log.Fatalln(err)
		}
	}

	// 以连接的当前偏移量读取消息，成功时将偏移量提前，以便对read方法下一次调用将生成下一条消息	 由于此方法为消息键和值分配内存缓冲区，因此内存效率低于读取，但具有io永不失败的优点
	if message, err := conn.ReadMessage(1024); err == nil {
		fmt.Println("读取内容对象：", message)
	} else {
		log.Fatalln(err)
	}

	// 从kafka服务器读取一批消息，该方法始终返回非nil的批处理(使用默认配置)
	batch := conn.ReadBatch(1024, 10240)
	// 关闭批处理
	defer func(batch *kafka.Batch) {
		if err = batch.Close(); err != nil {
			log.Fatalln(err)
		}
	}(batch)
	// 将批次中的下一条消息的值读取到b中，返回读取的字节数，如果无法读取下一条消息，则返回错误
	if n, err := batch.Read(rb); err == nil {
		fmt.Println("读取内容", string(rb), "读取字节数：", n)
	} else {
		log.Fatalln(err)
	}
	// 读取并返回批次中的下一条消息
	if message, err := batch.ReadMessage(); err == nil {
		fmt.Println("读取内容对象：", message)
	} else {
		log.Fatalln(err)
	}
	// 返回批处理分区
	fmt.Println("批处理分区：", batch.Partition())
	// 返回分区中的当前最高位
	fmt.Println("分区当前最高位", batch.HighWaterMark())
	// 提供kafka服务器在连接上应用的限制持续时间
	fmt.Println("限制持续时间：", batch.Throttle())
	// 如果批次中断，Err将返回非nil错误
	if err = batch.Err(); err != nil {
		log.Fatalln(err)
	}

	// 从kafka服务器读取一批消息，该方法始终返回非nil的批处理(自定义配置)
	batch = conn.ReadBatchWith(kafka.ReadBatchConfig{
		MinBytes:       1024,                  //	向代理指示使用者将接受的最小批大小
		MaxBytes:       10240,                 // 向代理指示使用者将接受的最大批大小
		IsolationLevel: kafka.ReadUncommitted, // 控制事务记录的可见性	ReadUncommitted：所有记录可见；ReadCommitted：只有非事务记录和已提交记录可见
		MaxWait:        0,                     // 代理等待达到最小/最大字节目标的时间量 此设置独立于任何网络级别超时或截止日期	当此字段为零时，将根据连接的读取截止日期推断最大等待时间
	})

	// 返回时间戳等于或大于 t 的第一条消息的偏移量
	offset, err = conn.ReadOffset(now)
	if err != nil {
		fmt.Println("偏移量：", offset)
	} else {
		log.Fatalln(err)
	}

	// 返回连接上可用的第一个偏移量
	if offset, err = conn.ReadFirstOffset(); err == nil {
		fmt.Println("第一个偏移量：", offset)
	} else {
		log.Fatalln(err)
	}

	// 返回连接上可用的最后偏移量
	if offset, err = conn.ReadLastOffset(); err == nil {
		fmt.Println("最后偏移量：", offset)
	} else {
		log.Fatalln(err)
	}

	// 返回连接使用的主题的第一个和最后一个绝对偏移量
	if firstOffsets, lastOffsets, err := conn.ReadOffsets(); err == nil {
		fmt.Println("第一个偏移量：", firstOffsets, "最后偏移量：", lastOffsets)
	} else {
		log.Fatalln(err)
	}

	// 返回给定主题列表的可用分区列表
	if partitions, err := conn.ReadPartitions(topic, topic2); err == nil {
		for _, partition := range partitions {
			fmt.Println("分区信息：", partition)
		}
	} else {
		log.Fatalln(err)
	}

	// 向kafka 代理写入建立此连接的消息。该方法返回写入的字节数，如果出现错误，则返回错误
	if n, err := conn.Write(wb); err == nil {
		fmt.Println("写入的字节数：", n)
	} else {
		log.Fatalln(err)
	}

	// 将一批消息写入连接的主题和分区，返回写入的字节数	写操作是一个原子操作，要么完全成功，要么失败	信息不压缩
	if n, err := conn.WriteMessages(m, m2); err == nil {
		fmt.Println("写入的字节数：", n)
	} else {
		log.Fatalln(err)
	}

	// 将一批消息写入连接的主题和分区，返回写入的字节数	写操作是一个原子操作，要么完全成功，要么失败	如果压缩编解码器不是nil，则消息将被压缩
	if n, err := conn.WriteCompressedMessages(&compress.GzipCodec, m, m2); err == nil {
		fmt.Println("写入的字节数：", n)
	} else {
		log.Fatalln(err)
	}

	// 将一批消息写入连接的主题和分区，返回写入的字节数、分区和偏移量以及kafka代理分配给消息集的时间戳	写操作是一个原子操作，要么完全成功，要么失败	如果压缩编解码器不是nil，则消息将被压缩
	if n, partition, offset, appendTime, err := conn.WriteCompressedMessagesAt(&compress.GzipCodec, m, m2); err == nil {
		fmt.Println("写入的字节数：", n, "分区信息：", partition, "偏移量：", offset, "代理分配给消息集的时间戳：", appendTime)
	} else {
		log.Fatalln(err)
	}

	// 设置连接在生成消息时请求的副本的确认数
	if err := conn.SetRequiredAcks(10); err != nil {
		log.Fatalln(err)
	}

}
