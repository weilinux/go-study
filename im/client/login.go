package main

import (
	"encoding/json"
	"fmt"
	"github.com/my/repo/im/common/message"
	"github.com/my/repo/im/common/utils"
	"net"
)

//TODO 登录

// login 校验登录
func login(uid *int, pwd *string) (err error) {
	//连接服务端
	conn, err := net.Dial("tcp", net.JoinHostPort("127.0.0.1", "8090"))
	if err != nil {
		return
	}

	//关闭连接
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	//创建消息工具结构体
	utils := utils.Transfer{
		Conn: conn,
	}

	//创建消息结构体
	messages := message.Message{
		//消息类型为登录消息
		Type: message.LoginMesType,
	}

	//创建消息主体结构体
	var loginMes message.LoginMes
	loginMes.UserId = *uid
	loginMes.UserPwd = *pwd
	//loginMes.UserName =

	//将消息主体进行序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		return
	}

	//将消息主体赋值给消息结构体
	messages.Data = string(data)

	//将消息结构体进行序列化
	mes, err := json.Marshal(messages)
	if err != nil {
		return
	}

	//发送信息
	err = utils.WriteMsg(mes)
	if err != nil {
		return
	}

	//获取信息
	msg, err := utils.ReadMsg()
	if err != nil {
		return
	}

	//将smg.Data反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(msg.Data), &loginResMes)
	if err != nil {
		return
	}

	//判断状态
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else {
		fmt.Println(loginResMes.Msg)
	}

	return nil
}
