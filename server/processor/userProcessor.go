package processor

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	Conn   net.Conn
	UserId int
}

//通知所有在线用户：userId通知其他在线用户自己上线了
func (this *UserProcessor) NotifyOthersMeOnline(userId int) {
	//遍历onlineUsers,一个个的发送NotifyUserStatusResMes
	for id, up := range onlineUsers.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyOthersMeOnlineDetails(userId)
	}

}

func (this *UserProcessor) NotifyOthersMeOnlineDetails(userId int) {
	//组装NotifyUserStatusResMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusResMesTyPe

	var notifyUserStatusResMes message.NotifyUserStatusResMes
	notifyUserStatusResMes.UserId = userId
	notifyUserStatusResMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusResMes)
	if err != nil {
		fmt.Println("marshal notifyUserStatusResMes failed in NotifyOthersMeOnlineDetails in function,err=", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal mes failed in NotifyOthersMeOnlineDetails in function,err=", err)
		return
	}

	tF := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tF.WritePkg(data)
	if err != nil {
		fmt.Println("Write pkg failed in function,err=", err)
		return
	}
	//return
}

//处理注册请求
func (this *UserProcessor) ServerProcessRegister(mes *message.Message) (err error) {
	//先从mes中取出mes.Data，并直接反序列化成LoginReqMes
	var registerReqMes message.RegisterReqMes
	err = json.Unmarshal([]byte(mes.Data), &registerReqMes)
	if err != nil {
		fmt.Println("unmarshal failed in ServerProcessRegister function,err=", err)
		return
	}

	//声明返回消息结构并赋值
	var registerResMes message.Message
	registerResMes.Type = message.RegisterResMesType
	//声明返回登录结果结构体并赋值
	var registerRes message.RegisterResMes
	err = model.MyUserDao.UserRegister(&registerReqMes.User)
	if err != nil {
		fmt.Println("user register failed in  ServerProcessRegister function,err=", err)
		if err == model.ERROR_USER_EXISTS {
			registerRes.Code = 111
			registerRes.Description = err.Error()
		} else {
			registerRes.Code = 222
			registerRes.Description = "服务器内部错误"
		}
		return
	} else {
		registerRes.Code = 200
		registerRes.Description = "注册成功"
	}

	//registerResMes.Data赋值
	data, err := json.Marshal(registerRes)
	if err != nil {
		fmt.Println("marshal registerRes failed in ServerProcessRegister,err=", err)
		return
	}
	registerResMes.Data = string(data)

	//对loginResMes序列化然后发送给客户端
	data, err = json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("marshal loginResMes failed in ServerProcessRegister,err=", err)
		return
	}
	tF := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tF.WritePkg(data)
	if err != nil {
		fmt.Println("write pkg failed in ServerProcessRegister function,err=", err)
		return
	}
	return

}

//处理登录请求
func (this *UserProcessor) ServerProcessLogin( /*conn net.Conn,*/ mes *message.Message) (err error) {
	//先从mes中取出mes.Data，并直接反序列化成LoginReqMes
	var loginReqMes message.LoginReqMes
	err = json.Unmarshal([]byte(mes.Data), &loginReqMes)
	if err != nil {
		fmt.Println("unmarshal failed in serverProcessLogin,err=", err)
		return
	}
	//声明返回消息结构并赋值
	var loginResMes message.Message
	loginResMes.Type = message.LoginResMesType
	//声明返回登录结果结构体并赋值
	var loginRes message.LoginResMes
	_, err = model.MyUserDao.LoginVerify(loginReqMes.UserId, loginReqMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_PWD {
			loginRes.Code = 403
			loginRes.Description = err.Error()
		} else if err == model.ERROR_USER_NOTEXISTS {
			loginRes.Code = 500
			loginRes.Description = err.Error()
		} else {
			loginRes.Code = 505
			loginRes.Description = "其他错误"
		}
	} else {
		loginRes.Code = 200
		loginRes.Description = "登陆成功"

		//登陆成功的用户应该放入onlineUsers
		this.UserId = loginReqMes.UserId
		onlineUsers.AddOnlineUsers(this)
		//通知其他在线用户，我上线了
		this.NotifyOthersMeOnline(loginReqMes.UserId)
		//将当前在线用户的id放入到loginResMes.usersId中
		for id, _ := range onlineUsers.onlineUsers {
			loginRes.UsersId = append(loginRes.UsersId, id)
		}

	}
	/*if loginReqMes.UserId == 100 && loginReqMes.UserPwd == "123456" {
		//用户合法登陆成功
		loginRes.Code = 200
		loginRes.Description = "登录成功！"
	} else {
		loginRes.Code = 500
		loginRes.Description = "登录失败！"
	}*/
	//将loginRes序列化再给loginResMes.Data赋值
	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("marshal loginRes failed in serverProcesssLogin,err=", err)
		return
	}
	loginResMes.Data = string(data)

	//对loginResMes序列化然后发送给客户端
	data, err = json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("marshal loginResMes failed in serverProcesssLogin,err=", err)
		return
	}
	tF := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tF.WritePkg(data)
	if err != nil {
		fmt.Println("write pkg failed in serverProcessLogin function,err=", err)
		return
	}
	return
}
