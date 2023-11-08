package models

import "time"

type tb_blog struct {
	ID         uint
	TagId      uint
	Title      string    `json:"title" gorm:"comment:'标题'"`
	Content    string    `json:"content" gorm:"comment:'内容'"`
	CreateTime time.Time `gorm:"comment:'创建时间';type:datetime;"`
	UpdateTime time.Time `gorm:"comment:'修改时间';type:datetime;"`
}
