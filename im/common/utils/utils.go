package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/my/repo/im/common/message"
	"net"
)

// Transfer 消息结构体
type Transfer struct {
	Conn net.Conn //连接结构体
	Buf  []byte   //数据缓冲
}

// ReadMsg 读取消息
func (t *Transfer) ReadMsg() (msg message.Message, err error) {
	//给消息结构体创建Buf切片的内存
	//小于等于4是因为发送消息只占用一个uint32(4个字节)，所有在这里使用会导致内存不够
	if cap(t.Buf) <= 4 {
		t.Buf = make([]byte, 1024)
	}

	//获取信息
	n, err := t.Conn.Read(t.Buf[:4])
	if n != 4 && err != nil {
		//自定义err
		//err = errors.New("read msg length fail")
		return
	}

	//根据buf[:4]切片转成uint32类型
	msgLen := binary.BigEndian.Uint32(t.Buf[:4])

	//根据msgLen读取信息内容 （信息长度为uint32占了前四个字节，所有这里取消息内容只需要从5开始取并取msgLen长度）
	n, err = t.Conn.Read(t.Buf[:msgLen])
	if n != int(msgLen) && err != nil {
		err = errors.New("read msg body fail")
		return
	}

	//将buf([]byte)转成message.Message结构体
	err = json.Unmarshal(t.Buf[:msgLen], &msg)
	if err != nil {
		return
	}

	return msg, nil
}

// WriteMsg 发送信息
func (t *Transfer) WriteMsg(msg []byte) (err error) {
	//获取消息长度
	mesLen := uint32(len(msg))

	//将消息长度（uint32）类型转成[]byte切片 (uint32 4*8 四个字节)
	if cap(t.Buf) == 0 {
		t.Buf = make([]byte, 4)
	}
	binary.BigEndian.PutUint32(t.Buf, mesLen)

	//发送消息长度
	n, err := t.Conn.Write(t.Buf)
	if n != 4 && err != nil {
		return err
	}
	//fmt.Println("消息长度发送成功 长度为：", len(msg))

	//发送消息本身
	n, err = t.Conn.Write(msg)
	if n != int(mesLen) && err != nil {
		return err
	}

	//fmt.Println("消息本身发送成功")
	return nil
}
