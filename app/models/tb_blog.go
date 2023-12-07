package models

type TbBlog struct {
	ID
	TagId   uint
	Title   string `json:"title" gorm:"comment:'标题'"`
	Content string `json:"content" gorm:"comment:'内容'"`
	Tag     string `json:"tag" gorm:"comment:'标签'"`
	Timestamps
}
