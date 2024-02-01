package models

type Banner struct {
	ID
	UID uint `form:"uid" json:"uid" gorm:"comment:用户ID"`
	MID uint `form:"mid" json:"mid" gorm:"comment:图片ID"`
	Timestamps
}
