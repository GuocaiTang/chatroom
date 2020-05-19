package processor

import (
	"chatroom/client/model"
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
}

func (this *UserProcessor) Register(userId int, userPwd, userName string) (err error) {
	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial failed in Register function,err=", err)
		return
	}
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterReqMesType
	//3.RegisterReqMes
	var registerReqMes message.RegisterReqMes
	registerReqMes.User.UserId = userId
	registerReqMes.User.UserPwd = userPwd
	registerReqMes.User.UserName = userName

	//4.registerReqMes
	registerReqData, err := json.Marshal(registerReqMes)
	if err != nil {
		fmt.Println("marshal register req mess failed in Register function,err=", err)
		return
	}
	//5.给mes.Data赋值
	mes.Data = string(registerReqData)

	//6.将mes序列化
	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal mes data failed in Register function,err=", err)
		return
	}

	//7.发送
	//fmt.Printf("客户端注册发送消息内容=%s,长度=%d \n", string(mesData), len(mesData))
	//发送消息
	tF := &utils.Transfer{
		Conn: conn,
	}
	err = tF.WritePkg(mesData)
	if err != nil {
		fmt.Println("write pkg failed in login in Register function,err=", err)
		return
	}

	//处理服务端返回的消息
	mes, err = tF.ReadPkg()
	if err != nil {
		fmt.Println("read pkg failed in Register function,err=", err)
		return
	}
	//将mes.Data反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println(registerResMes.Description)
		for {
			ShowLoginSuceessMenu()
		}
	} else {
		fmt.Println(registerResMes.Description)
	}
	return
}

func (this *UserProcessor) Login(userId int, userPwd string) (err error) {
	/*fmt.Printf("用户名：%d,密码：%s \n", userId, userPwd)
	return nil*/

	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial failed in Login function,err=", err)
		return
	}
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginReqMesType
	//3.创建一个LoginReqMes结构体
	var loginReqMes message.LoginReqMes
	loginReqMes.UserId = userId
	loginReqMes.UserPwd = userPwd

	//4.将loginReqMes序列化
	loginReqData, err := json.Marshal(loginReqMes)
	if err != nil {
		fmt.Println("marshal login req mess failed in Login function,err=", err)
		return
	}
	//5.给mes.Data赋值
	mes.Data = string(loginReqData)

	//6.将mes序列化
	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal mes data failed in Login function,err=", err)
		return
	}

	//7.mesData即为我们要发送的消息
	//7.1 先发送消息的长度
	//因为conn.Write()接收的是byte切片，所以需要将消息长度转成一个表示长度的byte切片
	/*var mesLen uint32
	mesLen = uint32(len(mesData))
	var mesLenByte [4]byte
	binary.BigEndian.PutUint32(mesLenByte[0:4], mesLen)
	//发送长度
	n, err := conn.Write(mesLenByte[0:4])
	if n != 4 || err != nil {
		fmt.Println("send len failed,err=", err)
		return
	}

	fmt.Printf("客户端发送消息内容=%s,长度=%d \n", string(mesData), len(mesData))

	//发送消息本身
	_,err = conn.Write(mesData)
	if err != nil {
		fmt.Println("conn write mesData failed,err=",err)
		return
	}*/
	//fmt.Printf("客户端登录发送消息内容=%s,长度=%d \n", string(mesData), len(mesData))
	//发送消息
	tF := &utils.Transfer{
		Conn: conn,
	}
	err = tF.WritePkg(mesData)
	if err != nil {
		fmt.Println("write pkg failed in login in Login function,err=", err)
		return
	}

	//休眠20秒
	/*time.Sleep(time.Second*20)
	fmt.Println("客户端休眠了20秒...")*/

	//处理服务端返回的消息
	mes, err = tF.ReadPkg()
	if err != nil {
		fmt.Println("read pkg failed in login function,err=", err)
		return
	}
	//将mes.Data反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		model.CurUser.Conn=conn
		model.CurUser.UserId=userId
		//fmt.Println(loginResMes.Description)
		//显示所有在线用户
		//fmt.Println("当前在线用户如下：")
		for _, v := range loginResMes.UsersId {
			if v == loginReqMes.UserId {
				fmt.Printf("------恭喜用户id:%d 登录成功！------\n", v)
				continue
			}
			//fmt.Println("用户id:\t", v)
			// 把在线用户保存到map
			/*user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user*/
			updateUserStatus(v, message.UserOnline)
		}
		//fmt.Print("\n\n")
		go serverKeep(conn)

		for {
			ShowLoginSuceessMenu()
		}
	} else {
		fmt.Println(loginResMes.Description)
	}
	return
}
