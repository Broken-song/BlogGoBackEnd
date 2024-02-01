package request

import "bloggobackend/app/models"

type Blog struct {
	ID       uint   `form:"id" json:"id" gorm:"comment:'博文ID'"`
	UID      uint   `form:"uid" json:"uid" gorm:"comment:'用户ID'"`
	UserName string `form:"user_name"json:"user_name" gorm:"comment:'用户名'"`
	Title    string `form:"title" json:"title" gorm:"comment:'标题'"`
	Content  string `form:"content" json:"content" gorm:"comment:'内容'"`
	Tag      string `form:"tag" json:"tag" gorm:"comment:'标签'"`
	models.Timestamps
}

func (blog Blog) GetMessages() ValidatorMessages {
	return ValidatorMessages{}
}
