package models

import "strconv"

type User struct {
	ID       uint
	Account  string `json:"name" gorm:"not null;comment:用户名"`
	Email    string `json:"email" gorm:"not null;index;comment:邮箱"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID))
}
