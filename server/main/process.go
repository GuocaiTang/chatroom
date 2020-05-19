package main

import (
	"chatroom/common/message"
	"chatroom/server/processor"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Process struct {
	Conn net.Conn
}

//根据客户端发送的消息种类，调用相应的函数处理
func (this *Process) ServerProcessMes( /*conn net.Conn,*/ mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginReqMesType:
		//处理登录
		userP := &processor.UserProcessor{
			Conn: this.Conn,
		}
		err = userP.ServerProcessLogin(mes)
	case message.RegisterReqMesType:
		//处理注册
		userP := &processor.UserProcessor{
			Conn: this.Conn,
		}
		err = userP.ServerProcessRegister(mes)
	case message.SmsMesType:
		//处理短消息
		smsP := &processor.SmsProcessor{}
		smsP.SendGroupMes(mes)
	case message.AloneSmsMesType:
		smsP := &processor.SmsProcessor{}
		smsP.SendAloneMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Process) processDetail() (err error) {
	//循环读取客户端发送的消息
	for {
		tF := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tF.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出...")
				return err
			} else {
				fmt.Println("read pkg failed,err=", err)
				return err
			}
		}
		fmt.Println("mes=", mes)
		/*userP := &processor.UserProcessor{
			Conn: this.Conn,
		}
		err = userP.ServerProcessLogin(&mes)*/
		err = this.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("server process mes failed in process,err=", err)
			return err
		}
	}
}
