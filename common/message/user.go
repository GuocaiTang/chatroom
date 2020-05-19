package message

type User struct {
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	UserPwd    string `json:"user_pwd"`
	UserStatus int    `json:"user_status"`
}
