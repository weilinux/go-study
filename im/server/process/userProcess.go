package process

import (
	"encoding/json"
	"github.com/my/repo/im/common/message"
	"github.com/my/repo/im/common/utils"
	"net"
)

// UserProcess 用户处理结构体
type UserProcess struct {
	Conn net.Conn //连接结构体
}

// ServerProcessLogin 处理登录请求
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//将mes.Data反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		return
	}

	//创建消息结构体
	resMes := message.Message{
		//消息类型为登录返回消息
		Type: message.LoginResMesType,
	}

	//创建消息主体结构体
	var loginResMes message.LoginResMes

	//验证用户
	if loginMes.UserId == 100 && loginMes.UserPwd == "qweasd" {
		//成功
		loginResMes.Code = message.Success
	} else {
		//失败
		loginResMes.Code = message.Error
		//描述
		loginResMes.Msg = "user auth fail"
	}

	//将loginResMes结构体序列化
	resData, err := json.Marshal(loginResMes)
	if err != nil {
		return
	}

	//将序列化后的loginResMes赋给resMes(消息主体)中
	resMes.Data = string(resData)

	//将resMes消息主体序列化
	msgData, err := json.Marshal(resMes)
	if err != nil {
		return
	}

	//创建消息工具结构体
	transfer := &utils.Transfer{
		Conn: u.Conn,
	}

	//发送给客户端
	err = transfer.WriteMsg(msgData)
	if err != nil {
		return
	}

	return nil
}
