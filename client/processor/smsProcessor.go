package processor

import (
	"chatroom/client/model"
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcessor struct {
}

func (this *SmsProcessor)ShowAloneMes(mes *message.Message){
	var alSmsMes message.AloneSmsMes
	if err := json.Unmarshal([]byte(mes.Data), &alSmsMes); err != nil {
		fmt.Println("unmarshal mes failed in ShowGroupMes function,err=", err)
		return
	}

	content := fmt.Sprintf("用户id：%d 对你说：%s", alSmsMes.UserId, alSmsMes.Content)
	fmt.Println(content)
	fmt.Println()
	fmt.Println()
}

func (this *SmsProcessor) ShowGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	if err := json.Unmarshal([]byte(mes.Data), &smsMes); err != nil {
		fmt.Println("unmarshal mes failed in ShowGroupMes function,err=", err)
		return
	}

	content := fmt.Sprintf("用户id：%d 对大家说：%s", smsMes.UserId, smsMes.Content)
	fmt.Println(content)
	fmt.Println()
	fmt.Println()
}

func (this *SmsProcessor) SendAloneMes(userId int, content string) {
	var mes message.Message
	mes.Type = message.AloneSmsMesType

	var alSmsMes message.AloneSmsMes
	alSmsMes.Content = content
	alSmsMes.RemoteUserId = userId         //目标用户
	alSmsMes.UserId = model.CurUser.UserId //发起用户

	data, err := json.Marshal(alSmsMes)
	if err != nil {
		fmt.Println("marshal alSmsMes failed in SendAloneMes function,err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err!= nil{
		fmt.Println("marshal mes failed in SendAloneMes function,err=",err)
		return
	}

	tF := &utils.Transfer{
		Conn: model.CurUser.Conn,
	}
	if err = tF.WritePkg(data); err != nil {
		fmt.Println("WritePkg failed in SednGroupMes function,err=", err)
		return
	}
	return
}

func (this *SmsProcessor) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = model.CurUser.UserId

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("marshal smsMes failed in SendGroupMes function,err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal mes failed in SendGroupMes function,err=", err)
		return
	}

	tF := &utils.Transfer{
		Conn: model.CurUser.Conn,
	}
	if err = tF.WritePkg(data); err != nil {
		fmt.Println("WritePkg failed in SednGroupMes function,err=", err)
		return
	}
	return
}
