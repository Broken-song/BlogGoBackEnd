package services

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/models"
	"bloggobackend/global"
	"errors"
)

type blogService struct {
}

var BlogService = new(blogService)

func (blogService *blogService) CreateBlog(params request.Blog) (err error, blog models.TbBlog) {
	blog = models.TbBlog{UID: params.UID, UserName: params.UserName, Title: params.Title, Content: params.Content, Tag: params.Tag}
	err = global.App.DB.Create(&blog).Error
	return
}

func (blogService *blogService) UpdateBlog(params request.Blog) (err error, blog models.TbBlog) {
	err = global.App.DB.First(&blog, params.ID).Error
	if err != nil {
		err = errors.New("博文不存在")
	}

	err = global.App.DB.Model(&blog).Updates(params).Error
	if err != nil {
		err = errors.New("更新数据时发生错误")
	}
	return
}

func (blogService *blogService) DeleteBlog(params request.Blog) (err error, blog models.TbBlog) {
	err = global.App.DB.First(&blog, params.ID).Error
	if err != nil {
		err = errors.New("博文不存在")
	}
	err = global.App.DB.Delete(&blog, params.ID).Error
	return
}
