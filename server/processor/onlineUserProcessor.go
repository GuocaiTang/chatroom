package processor

import (
	"fmt"
)

type OnlineUsers struct {
	onlineUsers map[int]*UserProcessor
}

//因为Onlineusers实例在服务器端有且只有一个
//频繁用到，将其定义为全局变量
var onlineUsers *OnlineUsers

func init() {
	onlineUsers = &OnlineUsers{
		onlineUsers: make(map[int]*UserProcessor, 1024),
	}
}

//对onlineUser进行添加
func (this *OnlineUsers) AddOnlineUsers(up *UserProcessor) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *OnlineUsers) DeleteOnlineUsers(up *UserProcessor) {
	delete(this.onlineUsers, up.UserId)
}

//返回当前所有在线用户
func (this *OnlineUsers) GetAllOnlineUsers() map[int]*UserProcessor {
	return this.onlineUsers
}

//根据userId返回
func (this *OnlineUsers) GetOnlineUserById(userId int) (up *UserProcessor, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok { //当前用户不在线
		err = fmt.Errorf("用户%d不在线", userId)
		return
	}
	return
}
