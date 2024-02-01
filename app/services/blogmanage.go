package services

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/models"
	"bloggobackend/global"
	"strconv"
)

type blogManageService struct {
}

var BlogManageService = new(blogManageService)

type BlogManageResponse struct {
	Count int64           `json:"count"`
	Data  []models.TbBlog `json:"data"`
	Code  uint            `json:"code"`
}

func (blogmanage *blogManageService) GetBlogByPage(params request.BlogManage) BlogManageResponse {
	page := params.Page
	limit := params.Limit
	uid := params.Uid
	var tbblog []models.TbBlog
	err, user := UserService.GetUserInfo(strconv.Itoa(int(uid)))
	if err != nil {
	}
	if user.Role == 1 {
		global.App.DB.Raw("SELECT * FROM tb_blogs WHERE deleted_at IS NULL AND uid = ? ORDER BY updated_at DESC LIMIT ? OFFSET ?", uid, limit, (page-1)*limit).Find(&tbblog)
	} else if user.Role == 2 {
		global.App.DB.Raw("SELECT * FROM tb_blogs WHERE deleted_at IS NULL ORDER BY updated_at DESC LIMIT ? OFFSET ?", limit, (page-1)*limit).Find(&tbblog)
	}
	var count int64
	global.App.DB.Model(&models.TbBlog{}).Count(&count)
	data := BlogManageResponse{
		Count: count,
		Data:  tbblog,
		Code:  0,
	}
	return data
}
