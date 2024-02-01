package request

type BannerGet struct {
	Business string `form:"business" json:"business" binding:"required"`
	UID      uint   `form:"uid" json:"uid" binding:"required" gorm:"comment:用户ID" `
}

func (bannerGet BannerGet) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"business.required": "业务类型不能为空",
		"uid.required":      "请确认登录状态",
	}
}
