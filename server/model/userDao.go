package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var MyUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

//工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	return &UserDao{
		pool: pool,
	}
}

//根据用户id返回一个用户实例以及错误信息
func (this *UserDao) getUserById(conn redis.Conn, userId int) (user *User, err error) {
	//根据id到redis查询用户信息
	res, err := redis.String(conn.Do("HGet", "users", userId))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//将res反序列化为user实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("unmarshal res to user failed in getUserById function,err=", err)
		return
	}
	return
}

//完成登陆的校验
//1.用户的id和pwd都正确，返回一个user实例、
//2.若用户名不存在或密码不正确则返回对应错误
func (this *UserDao) LoginVerify(userId int, userPwd string) (user *User, err error) {
	//先从UserDao的连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		fmt.Println("get user instant failed in LoginVerify function,err=", err)
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

//处理用户注册
func (this *UserDao) UserRegister(user *message.User) (err error) {
	//先从UserDao的连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()

	//校验用户是否已存在
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("marshal user failed in UserRegister function,err=", err)
		return
	}
	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("hset user to db failed in UserRegister function,err=", err)
		return
	}
	return
}
