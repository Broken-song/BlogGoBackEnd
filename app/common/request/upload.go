package request

import "mime/multipart"

type ImageUpload struct {
	Business string                `form:"business" json:"business" binding:"required"`
	UID      uint                  `form:"uid" json:"uid" gorm:"comment:用户ID"`
	Image    *multipart.FileHeader `form:"image" json:"image" binding:"required"`
}

func (imageUpload ImageUpload) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"business.required": "业务类型不能为空",
		"image.required":    "请选择图片",
	}
}
