package main

import (
	"fmt"
	"github.com/my/repo/im/common/message"
	"github.com/my/repo/im/common/utils"
	"github.com/my/repo/im/server/process"
	"io"
	"net"
)

// Processor 消息分发结构体
type Processor struct {
	Conn net.Conn //连接结构体
}

// 根据客户端发送消息种类不同，决定调用哪个函数来处理
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	//处理登录
	case message.LoginMesType:
		//创建用户处理结构体
		up := &process.UserProcess{
			Conn: p.Conn,
		}

		//用户处理
		err := up.ServerProcessLogin(mes)
		if err != nil {
			return err
		}
	default:
		fmt.Println("消息类型错误")

	}

	return nil
}

// ServerProcessHandle 消息总控
func (p *Processor) ServerProcessHandle() (err error) {
	//循环读取客户端发送的信息
	for {
		//创建消息结构体
		transfer := utils.Transfer{
			Conn: p.Conn,
		}

		//获取信息
		mes, err := transfer.ReadMsg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端断开连接...")
			} else {
				fmt.Println(err)
			}
			return err
		}

		//处理消息
		err = p.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
