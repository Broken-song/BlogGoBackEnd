package services

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/models"
	"bloggobackend/global"
)

type userManageService struct {
}

var UserManageService = new(userManageService)

type UserManageResponse struct {
	Count int64         `json:"count"`
	Data  []models.User `json:"data"`
	Code  uint          `json:"code"`
}

func (usermanage *userManageService) GetUserByPage(params request.UserManage) UserManageResponse {
	page := params.Page
	limit := params.Limit
	var user []models.User
	global.App.DB.Raw("SELECT id, account, email, created_at, updated_at, content, role FROM users WHERE deleted_at IS NULL ORDER BY updated_at DESC LIMIT ? OFFSET ?", limit, (page-1)*limit).Find(&user)
	var count int64
	global.App.DB.Model(&models.User{}).Count(&count)
	data := UserManageResponse{
		Count: count,
		Data:  user,
		Code:  0,
	}
	return data
}
