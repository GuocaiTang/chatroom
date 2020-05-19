package model

import (
	"chatroom/common/message"
	"net"
)

//可声明为全局，在登陆成功后初始化
var CurUser CurrentUser

type CurrentUser struct {
	Conn net.Conn
	message.User
}
