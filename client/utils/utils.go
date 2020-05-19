package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg( /*conn net.Conn*/) (mes message.Message, err error) {
	//fmt.Println("读取客户端发送的数据...")
	//buf := make([]byte, 8096)
	//conn.Read在没有被关闭的情况下才会阻塞
	//若客户端关闭了conn，则不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//fmt.Println("conn.Read(buf[:4]) failed,err=", err)
		return
	}
	//fmt.Println("读取到的buff=",buf[:4])
	//根据buf[:4]转换成一个uint32
	var mesDataLen uint32
	mesDataLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据mesDataLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:mesDataLen])
	if n != int(mesDataLen) || err != nil {
		//fmt.Println("conn.Read(buf[:mesDataLen]) failed,err=", err)
		return
	}
	//buf[:mesDataLen]反序列化成Message
	err = json.Unmarshal(this.Buf[:mesDataLen], &mes)
	if err != nil {
		fmt.Println("unmarshal buf failed in readPkg,err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg( /*conn net.Conn, */ data []byte) (err error) {
	//获取数据的长度切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("send pkg len failed in writePkg function,err=", err)
		return
	}
	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("send pkg data failed in writePkg function,err=", err)
		return
	}
	return
}
