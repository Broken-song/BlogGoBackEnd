package services

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/models"
	"bloggobackend/global"
)

type dataService struct {
}

var DataService = new(dataService)

type DataResponse struct {
	TotalCount int64           `json:"total_count"`
	PageCount  int             `json:"page_count"`
	Content    []models.TbBlog `json:"content"`
}

func (dataService *dataService) GetDataByPage(params request.DataGet) DataResponse {
	page := params.Page
	limit := params.Limit
	order := params.Order
	var tbblog []models.TbBlog
	if order == 0 {
		//global.App.DB.Order("updated_at DESC").Limit(int(limit)).Offset(int((page - 1) * limit)).Find(&tbblog)
		global.App.DB.Raw("SELECT * FROM tb_blogs WHERE deleted_at IS NULL ORDER BY updated_at DESC LIMIT ? OFFSET ?", limit, (page-1)*limit).Find(&tbblog)
	} else if order == 1 {
		//global.App.DB.Find(&tbblog, "ORDER BY updated_at ASC LIMIT ? OFFSET ?", limit, (page-1)*limit)
		global.App.DB.Raw("SELECT * FROM tb_blogs WHERE deleted_at IS NULL ORDER BY updated_at ASC LIMIT ? OFFSET ?", limit, (page-1)*limit).Find(&tbblog)
	}
	var count int64
	global.App.DB.Model(&models.TbBlog{}).Count(&count)
	data := DataResponse{
		TotalCount: count,
		PageCount:  len(tbblog),
		Content:    tbblog,
	}
	return data
}
