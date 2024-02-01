package request

type UserManage struct {
	Page  uint `form:"page" json:"page" binding:"required"`
	Limit uint `form:"limit" json:"limit" binding:"required"`
}

func (usermanage UserManage) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"page.required":  "当前页码不能为空",
		"limit.required": "请确认每页需要显示的条目数",
	}
}
