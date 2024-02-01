package models

type TbBlog struct {
	ID
	UID      uint   `json:"uid" gorm:"comment:'用户ID'"`
	UserName string `json:"user_name" gorm:"comment:'用户名'"`
	Title    string `json:"title" gorm:"comment:'标题'"`
	Content  string `json:"content" gorm:"comment:'内容'"`
	Tag      string `json:"tag" gorm:"comment:'标签'"`
	Timestamps
	SoftDeletes
}
