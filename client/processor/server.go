package processor

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowLoginSuceessMenu() /*(err error)*/ {
	fmt.Println("你可以进行如下操作：")
	fmt.Println("------1.显示在线用户列表")
	fmt.Println("------2.群发送消息")
	fmt.Println("------3.私聊")
	fmt.Println("------4.退出系统")
	fmt.Println("请选择（1-4）：")
	var userChoosen int
	fmt.Scanln(&userChoosen)
	var sP *SmsProcessor
	switch userChoosen {
	case 1:
		//fmt.Println("显示在线用户列表")
		showOnlineUsers()
	case 2:
		fmt.Println("请输入您要群发的消息：")
		var content string
		fmt.Scanln(&content)
		sP.SendGroupMes(content)
	case 3:
		//fmt.Println("信息列表")
		fmt.Println("请输入对方用户id:")
		var userId int
		fmt.Scanln(&userId)
		fmt.Println("请输入您要发送的消息：")
		var content string
		fmt.Scanln(&content)
		sP.SendAloneMes(userId, content)
	case 4:
		fmt.Println("退出系统...")
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重新输入...")
	}
}

//和服务器保持通讯
func serverKeep(conn net.Conn) /*(err error)*/ {
	tF := &utils.Transfer{
		Conn: conn,
	}
	for {
		//fmt.Println("客户端私下在等待读取服务端的消息...")
		mes, err := tF.ReadPkg()
		if err != nil {
			fmt.Println("read pkg from server failed in serverKeep function,err=", err)
			return
		}
		//读取到消息进行下一步处理
		//fmt.Println("mes=", mes)
		switch mes.Type {
		case message.NotifyUserStatusResMesTyPe:
			//取出消息反序列化
			var notifyUserStatusResMes message.NotifyUserStatusResMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusResMes)
			//保存到客户端的onlineUsers中
			updateUserStatus(notifyUserStatusResMes.UserId, notifyUserStatusResMes.Status)
		case message.SmsMesType:
			smsP := SmsProcessor{}
			smsP.ShowGroupMes(&mes)
		case message.AloneSmsMesType:
			smsP := SmsProcessor{}
			smsP.ShowAloneMes(&mes)
		default:
			fmt.Println("未知消息类型")
		}
	}
}
