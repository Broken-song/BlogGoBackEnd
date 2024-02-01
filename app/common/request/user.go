package request

import (
	"bloggobackend/app/models"
	"github.com/dgrijalva/jwt-go"
)

type Register struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 自定义错误信息
func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"account.required":  "用户名不能为空",
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式错误",
		"password.required": "密码不能为空",
	}
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式不正确",
		"password.required": "密码不能为空",
	}
}

type ProfileSet struct {
	UID      uint   `json:"uid" binding:"required"`
	Account  string `json:"account"`
	Email    string `json:"email"`
	Content  string `json:"content"`
	CurrPass string `json:"current_password"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
	models.Timestamps
}

func (profileSet ProfileSet) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"uid.required": "请检查登录状态",
	}
}

type User struct {
	UID   uint `json:"uid" binding:"required"`
	Token *jwt.Token
	models.Timestamps
}

func (user User) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"uid.required": "请检查登录状态",
	}
}
