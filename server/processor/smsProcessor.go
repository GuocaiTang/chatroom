package processor

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcessor struct {
}

func (this *SmsProcessor) SendAloneMes(mes *message.Message) {
	var alSmsMes message.AloneSmsMes
	if err := json.Unmarshal([]byte(mes.Data), &alSmsMes); err != nil {
		fmt.Println("unmarsal mesData failed in SendAloneMes function,err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal smsMes failed,err=", err)
		return
	}

	for id, up := range onlineUsers.onlineUsers {
		if id == alSmsMes.RemoteUserId {
			this.SendAloneMesDetail(data, up.Conn)
		}
	}

}

func (this *SmsProcessor) SendAloneMesDetail(data []byte, conn net.Conn) {
	tF := &utils.Transfer{
		Conn: conn,
	}
	if err := tF.WritePkg(data); err != nil {
		fmt.Println("write pkg failed,err=", err)
		return
	}
	return
}
func (this *SmsProcessor) SendGroupMes(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	if err = json.Unmarshal([]byte(mes.Data), &smsMes); err != nil {
		fmt.Println("Unmarshal mesData failed in SendGroupMes function,err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal smsMes failed,err=", err)
		return
	}

	for id, up := range onlineUsers.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendGroupMesToEachOnlineUsers(data, up.Conn)
	}
	return
}

func (this *SmsProcessor) SendGroupMesToEachOnlineUsers(data []byte, conn net.Conn) {
	tF := &utils.Transfer{
		Conn: conn,
	}
	if err := tF.WritePkg(data); err != nil {
		fmt.Println("write pkg failed,err=", err)
		return
	}
	return
}
