package processor

import (
	"chatroom/common/message"
	"fmt"
)

//客户端要维护的map,在登录成功出初始化
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

//客户端显示当前在线用户
func showOnlineUsers() {
	//即遍历onlineUsers
	fmt.Println("当前在线用户如下：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
	fmt.Print("\n\n")
}

//处理返回的NotifyUserStatusMes
func updateUserStatus(userId int, status int) {
	user, ok := onlineUsers[userId]
	if !ok {
		user = &message.User{
			UserId: userId,
		}
	}
	user.UserStatus = status
	onlineUsers[userId] = user

	//showOnlineUsers()
}
