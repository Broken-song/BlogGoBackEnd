package request

type DataGet struct {
	Page  uint `form:"page" json:"page" binding:"required"`
	Limit uint `form:"limit" json:"limit" binding:"required"`
	Order uint `form:"order" json:"order"` // 排序方式（升序/降序）
}

func (dataget DataGet) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"page.required":  "当前页码不能为空",
		"limit.required": "请确认每页需要显示的条目数",
	}
}
