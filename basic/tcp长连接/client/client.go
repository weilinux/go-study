package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//关闭客户端连接
func closeClient(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Printf("客户端%s关闭失败：%s\n", conn.RemoteAddr(), err)
		return
	}

	fmt.Printf("客户端关闭连接")
}

func main() {
	//TODO TCP连接

	//地址、端口
	host, port := "127.0.0.1", "8888"

	//连接tcp服务
	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		fmt.Println("client connect error：", err)
		return
	}
	fmt.Println("client connect success：", conn)

	//实例reader缓存
	for {
		reader := bufio.NewReader(os.Stdin)

		//从命令行每行读取信息
		line, b, err := reader.ReadLine()
		if err != nil {
			fmt.Println("reader line error：", err)
			return
		}
		fmt.Println("reader line isPrefix：", b)

		if string(line) == "exit" {
			go closeClient(conn)
			return
		}

		//客户端发送消息给服务
		write, err := conn.Write(line)
		if err != nil {
			fmt.Println("connect write info error：", err)
			return
		}
		fmt.Println("connect write info success：", write)
	}
}
