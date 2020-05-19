package message

const (
	LoginReqMesType            = "LoginReqMes"
	LoginResMesType            = "LoginResMes"
	RegisterReqMesType         = "RegisterReqMes"
	RegisterResMesType         = "RegisterResMes"
	NotifyUserStatusResMesTyPe = "NotifyUserStatusResMes"
	SmsMesType                 = "SmsMes"
	AloneSmsMesType            = "AloneSmsMes"
)
const (
	UserOnline = iota
	UserOffline
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息实体
}

type LoginReqMes struct {
	UserId   int    `json:"user_id"`   //用户ID
	UserPwd  string `json:"user_pwd"`  //用户密码
	UserName string `json:"user_name"` //用户名
}

type LoginResMes struct {
	Code        int    `json:"code"` //状态码：500.用户未注册  200.登陆成功
	UsersId     []int  `json:"users_id"`
	Description string `json:"description"` //返回描述
}

type RegisterReqMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type NotifyUserStatusResMes struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

//发送的短消息
type SmsMes struct {
	Content string `json:"content"`
	User
}

type AloneSmsMes struct {
	Content      string `json:"content"`
	RemoteUserId int    `json:"remote_user_id"`
	User
}
