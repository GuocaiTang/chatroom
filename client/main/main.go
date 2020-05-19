package main

import (
	"chatroom/client/processor"
	"fmt"
)

var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择变量
	var userChoosen int
	//退出循环变量
	var loop bool = true

	for loop {
		fmt.Println("----------欢迎登陆聊天系统---------")
		fmt.Println("\t\t\t  1.用户登录")
		fmt.Println("\t\t\t  2.用户注册")
		fmt.Println("\t\t\t  3.退出系统")

		fmt.Println("请选择1-3进行操作：")

		fmt.Scanln(&userChoosen)
		switch userChoosen {
		case 1:
			fmt.Println("即将登陆...")
			fmt.Println("请输入用户ID：")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanln(&userPwd)
			userPro := &processor.UserProcessor{
			}
			err := userPro.Login(userId, userPwd)
			if err != nil {
				fmt.Println("user login failed in client main function,err=", err)
				return
			}
			//loop = false
		case 2:
			fmt.Println("即将注册...")
			//loop = false
			fmt.Println("请输入用户ID：")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入用户名：")
			fmt.Scanln(&userName)
			userPro := &processor.UserProcessor{
			}
			err := userPro.Register(userId, userPwd,userName)
			if err != nil {
				fmt.Println("user register failed in client main function,err=", err)
				return
			}
		case 3:
			fmt.Println("即将退出系统...")
			loop = false
		default:
			fmt.Println("用户输入有误，请重新输入...")
		}
	}
}
