package main

import (
	"fmt"
	"net"
)

//处理和客户端的通讯
func handle(conn net.Conn) {
	//关闭连接
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	//创建消息分发结构体
	pc := &Processor{
		Conn: conn,
	}

	//调用消息总控进行消息分发
	err := pc.ServerProcessHandle()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	//TODO tcp服务

	//监听
	fmt.Println("服务器在8090端口监听...")
	listen, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", "8090"))
	if err != nil {
		panic(err)
	}

	//关闭服务
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			panic(err)
		}
	}(listen)

	//等待客户端连接
	for {
		fmt.Println("等待客户端连接...")

		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		//用协程处理消息
		go handle(conn)
	}
}
