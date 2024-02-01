package models

import "strconv"

type User struct {
	ID       uint
	Account  string `json:"name" gorm:"not null; comment:用户名"`
	Email    string `json:"email" gorm:"not null;index; comment:邮箱"`
	Password string `json:"password" gorm:"not null;default:''; comment:用户密码"`
	Content  string `json:"content" gorm:"default:'这是一个很懒的人'; comment:用户简介"`
	Role     uint   `json:"role" gorm:"default:1; comment:用户权限级别"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID))
}
