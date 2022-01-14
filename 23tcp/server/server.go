package main

import (
	"fmt"
	"io"
	"net"
)

//接收客户端发送的信息
func process(conn net.Conn) {
	//关闭客户端
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Printf("客户端%s关闭失败：%s\n", conn.RemoteAddr(), err)
			return
		}

		fmt.Print("客户端关闭连接\n")
	}()

	//var dataBuffer bytes.Buffer

	for {
		//创建byte切片用来接收信息
		buf := make([]byte, 1024)

		//读取发送的信息
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端%s退出\n", conn.RemoteAddr())
			} else {
				fmt.Printf("客户端%s信息读取失败：%v\n", conn.RemoteAddr(), n)
			}
			return
		}

		fmt.Printf("客户端%s输出消息：%s\n", conn.RemoteAddr(), buf[:n])
		//dataBuffer.Write(buf[:read])
	}

	//fmt.Printf("客户端%s输出消息：%s\n", conn.RemoteAddr(), dataBuffer.String())
}

//获取客户端连接数
/*func monitor(listen net.Listener) {
	//获取当前客户端连接数
	listen.
}*/

func main() {
	//TODO TCP服务

	//地址、端口
	host, port := "0.0.0.0", "8888"

	//network表示监视类型，address表示监听地址 监听类型：tcp，地址：0.0.0.0，端口：8888
	listen, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		fmt.Println("listen error：", err)
		return
	}
	fmt.Println("listen suc：", listen)

	//关闭服务
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println("listen close error：", err)
			return
		}
		fmt.Println("listen close success")
	}(listen)

	//持续阻塞主函数结束，不停的接收客户端连接
	for {
		//等待客户端连接并阻塞主函数结束运行
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error：", err)
			return
		}
		fmt.Println("客户端连接地址：", conn.RemoteAddr())

		//通过携程循环接收客户端发送的信息
		go process(conn)
	}
}
