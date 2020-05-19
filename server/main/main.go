package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)

//对UserDao初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func init() {
	//当服务启动时，就初始化redis的连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	//初始化UserDao
	initUserDao()
}

//处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()
	//循环读取客户端发送的消息
	pc := &Process{
		Conn: conn,
	}
	err := pc.processDetail()
	if err != nil {
		fmt.Println("协程通讯错误，err=", err)
		return
	}
}

func main() {
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net listen failed,err=", err)
		return
	}
	defer listen.Close()

	//一旦监听成功，等待客户端来链接服务器
	fmt.Println("等待客户端来链接服务器...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept failed,err=", err)
		}
		//一旦链接成功，则启动协程和客户端保持通讯
		go process(conn)
	}
}
