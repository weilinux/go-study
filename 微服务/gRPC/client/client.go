package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/my/repo/微服务/gRPC/router"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func runFirst(client router.RouteGuideClient) {
	// 调用服务端GetFeature方法
	feature, err := client.GetFeature(ctx, &router.Point{
		Latitude:  310235000,
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

func runSecond(client router.RouteGuideClient) {
	// 调用服务端ListFeatures方法
	stream, err := client.ListFeatures(ctx, &router.Rectangle{
		Lo: &router.Point{
			Latitude:  313374060,
			Longitude: 121358540,
		},
		Hi: &router.Point{
			Latitude:  311034130,
			Longitude: 121598790,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// 等待服务端发来的流式响应(阻塞等待状态)
		if feature, err := stream.Recv(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println(feature)
		}
	}
}

func runThird(client router.RouteGuideClient) {
	// 调用服务端RecordRoute方法
	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// dummy data
	points := []*router.Point{
		{Latitude: 313374060, Longitude: 121358540},
		{Latitude: 311034130, Longitude: 121598790},
		{Latitude: 310235000, Longitude: 121437403},
	}

	for _, point := range points {
		// 使用stream将数据发送给服务端
		if err := stream.Send(point); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}

	// 关闭流并获取返回的消息
	summary, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(summary)
}

func readIntFromCommandLine(reader *bufio.Reader, target *int32) {
	_, err := fmt.Fscanf(reader, "%d\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func runForth(client router.RouteGuideClient) {
	// 调用服务端Recommend方法
	stream, err := client.Recommend(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// this goroutine listen to the server stream
	go func() {
		// 等待服务端发来的流式响应(阻塞等待状态)
		feature, err2 := stream.Recv()
		if err2 != nil {
			log.Fatalln(err2)
		}

		fmt.Println("Recommended：", feature)
	}()

	// 创建一个读缓冲区，使用标准输入
	reader := bufio.NewReader(os.Stdin)

	for {
		// 客户端请求体（通过cmd输入记录经纬坐标和计算模式）
		request := router.RecommendationRequest{Point: new(router.Point)}
		// 计算模式 0：最远的 1：最近的
		var mode int32
		fmt.Println("Enter Recommendation Mode (0 for farthest, 1 for the nearest)：")
		readIntFromCommandLine(reader, &mode)
		fmt.Println("Enter Latitude：")
		readIntFromCommandLine(reader, &request.Point.Latitude)
		fmt.Println("Enter Longitude：")
		readIntFromCommandLine(reader, &request.Point.Longitude)
		// 将计算模式赋值请求体
		request.Mode = router.RecommendationMode(mode)
		// 使用stream将数据发送给服务端
		if err := stream.Send(&request); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

var ctx = context.Background()

func main() {
	// 客户端创建连接
	conn, err := grpc.Dial(net.JoinHostPort("127.0.0.1", "5000"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	}

	// 程序退出前关闭连接
	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Fatalln(err)
		}
	}(conn)

	// 将连接传给routeGuideClient并返回（grpc已实现所有客户端方法）
	client := router.NewRouteGuideClient(conn)

	// 测试方法1
	runFirst(client)

	// 测试方法2
	runSecond(client)

	// 测试方法3
	runThird(client)

	// 测试方法4
	runForth(client)
}
