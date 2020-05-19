# 项目需求
----------欢迎登陆聊天系统---------  

                          1.用户登录
                          2.用户注册
                          3.退出系统
请选择1-3进行操作：  

------恭喜用户id:100 登录成功！------  
你可以进行如下操作：  
 ------1.显示在线用户列表  
 ------2.群发送消息  
 ------3.私聊  
 ------4.退出系统  
 请选择（1-4）：  
 
# 项目结构
见image
# 用户登陆流程分析（举例）
## 客户端
1.接收输入的userId和userPwd并且向服务端发送  
2.接收服务端返回的结果  
3.判断成功失败，显示相应的页面  
### 关键 设计消息结构
type Message struct{  
    Type string  //消息类型：登录消息/登出消息  
    Data string  //序列化后的具体消息
}  

type LoginReqMes struct{  
    userId int  
    userPwd string  
}  

### 数据发送流程  
1.实例化一个Message结构体  
2.Message.Type=登录消息类型，Message.Data=序列化后的LoginReqMes  
type LoginResMes struct{      
code int      
error string  
}  
3.序列化Message  
4.避免丢包  
（1）先向服务器发送Message的字节长度N  
（2）再发送消息本身  

## 服务端
1.接受客户端的userId和userPwd[goroutine]  
2.登录操作[比较]  
3.返回结果  
### 数据接收流程  
1.接收客户端发送的Message字节长度  
2.根据接收到的字节长度再接收消息本身  
3.接收时要判断实际接收到的消息内容长度是否等于客户端发送的字节长度  
4.如果不相等，启用纠错协议  
5.若想等：取到的消息->反序列化->Message  
6.取出Message.Data->反序列化->LonginReqMes  
7.取出LoginReqMes.userId和LoginReqMes.userPwd  
8.比较  
9.组装LoginResMes,返回客户端  
type LoginResMes struct{      
code int      
error string  
}  