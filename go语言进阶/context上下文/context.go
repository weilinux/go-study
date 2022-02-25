package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type context1 string

type context2 string

var wg = sync.WaitGroup{}

func f2(ctx context.Context) {
	// 执行完毕后关闭上下文
	defer wg.Done()

	for {
		//阻塞一秒
		time.Sleep(1 * time.Second)

		select {
		//监听上下文是否关闭  channel管道类型
		case <-ctx.Done():
			fmt.Println("ctx bone2")
			return
		default:
			fmt.Println("wait...2")
		}
	}
}

func f1(ctx context.Context) {
	// 执行完毕后关闭上下文
	defer wg.Done()

	//阻塞一秒
	wg.Add(1)

	//创建继承自该父级节点的子节点上下文
	ctx2, cancel2 := context.WithCancel(ctx)

	//添加协程执行数量
	go f2(ctx2)

	for {
		time.Sleep(1 * time.Second)

		select {
		//监听上下文是否关闭  channel管道类型
		case <-ctx.Done():
			fmt.Println("ctx bone1")

			//关闭子级上下文
			cancel2()
			return
		default:
			fmt.Println("wait...1")
		}
	}
}

func main() {
	// Background通常被用于主函数、初始化以及测试中，作为一个顶层的context，也就是说一般我们创建的context都是基于Background
	//context.Background()

	// TODO是在不确定使用什么context的时候才会使用
	//context.TODO()

	// 创建一个上下文 返回一个context类型的值和关闭该值的方法 无过期时间
	ctx, cancel := context.WithCancel(context.Background())

	// 创建一个上下文,需设置有效期(时间类型，格式为：time.Time) 返回一个context类型的值和关闭该值的方法 有过期时间
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))

	// 创建一个上下文,需设置多长时间后过期(时间间隔类型，格式为：time.Duration) 返回一个context类型的值和关闭该值的方法 有过期时间
	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	// 创建一个上下文,需设置键和值(建议键名已type的方式单独定义，这样在递归查询时可避免键名重复) 返回一个context类型的值 无过期时间 一般用于当作子级上下文使用
	var key1 context1 = "context1"
	//ctx := context.WithValue(context.Background(), key1, "value1")

	// 判断该通道是否已经关闭 (上下文是否已关闭 cancel())
	<-ctx.Done()

	// 获取上下文指定类型键的值,如果该上下文具有子级则会递归查询 默认值：nil
	fmt.Println(ctx.Value(key1))

	// 获取上下文的过期时间，只有WithDeadline、WithTimeout才具有 默认值：nil
	fmt.Println(ctx.Deadline())

	// 获取上下文的是否错误 默认值：nil
	fmt.Println(ctx.Err())

	// 错误类型：
	//context.Canceled 当前上下文被关闭返回该错误  类型：string

	//context.DeadlineExceeded 当前上下文超时返回该错误 类型：结构体
	//context.DeadlineExceeded.Error()  返回错误详情
	//context.DeadlineExceeded.Timeout() 判断是否超时
	//context.DeadlineExceeded.Temporary() 判断是否为临时的上下文(具有过期时间)

	// 关闭上下文，如果该上下文为父级则继承自该父级节点的上下文都将会关闭
	ctx.Done()

	// 添加协程执行数量
	wg.Add(1)

	// 启动协程
	go f1(ctx)

	// 阻塞五秒
	time.Sleep(5 * time.Second)

	// 关闭上下文
	cancel()

	// 协程未执行完成则阻塞主函数结束运行
	wg.Wait()
}
