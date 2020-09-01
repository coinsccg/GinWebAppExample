package models

import (
	"encoding/json"
	"errors"
)

// 定义结构体信息
type SignUpForm struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// UnmarshalJSON SignUpForm类型 自定义JSON反序列化的方法
func (s *SignUpForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return err
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else if required.Password != required.ConfirmPassword {
		err = errors.New("两次密码不一致")
	} else {
		s.UserName = required.UserName
		s.Password = required.Password
		s.ConfirmPassword = required.ConfirmPassword
	}
	return
}

type User struct {
	UserID   uint64 `json:"user_id" db:"user_id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
