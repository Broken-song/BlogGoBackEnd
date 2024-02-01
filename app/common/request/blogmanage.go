package request

type BlogManage struct {
	Page  uint `form:"page" json:"page"`
	Limit uint `form:"limit" json:"limit"`
	Uid   uint `form:"uid" json:"uid" binding:"required"`
}

func (blogmanage BlogManage) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"page.required":  "当前页码不能为空",
		"limit.required": "请确认每页需要显示的条目数",
	}
}
