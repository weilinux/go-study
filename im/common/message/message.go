package message

//TODO 消息

const (
	Success         = 200           //服务端返回code码  成功
	Error           = 500           //服务端返回code码  失败
	LoginMesType    = "LoginMes"    //客户端发送登录类型
	LoginResMesType = "LoginResMes" //服务端返回登录类型
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息主体
}

type LoginMes struct {
	UserId   int    `json:"user_id"`   //用户id
	UserPwd  string `json:"user_pwd"`  //用户密码
	UserName string `json:"user_name"` //用户名
}

type LoginResMes struct {
	Code int    `json:"code"` //状态码
	Msg  string `json:"msg"`  //错误信息
}
