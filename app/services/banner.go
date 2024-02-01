package services

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/models"
	"bloggobackend/global"
	"context"
	"strconv"
)

type bannerService struct {
}

var BannerService = new(bannerService)

func (bannerService *bannerService) GetUrlByID(params request.BannerGet) string {
	id := params.UID
	if id == 0 {
		return ""
	}

	var url string
	cacheKey := mediaCacheKeyPre + strconv.FormatUint(uint64(id), 10)

	exist := global.App.Redis.Exists(context.Background(), cacheKey).Val()
	if exist == 1 {
		url = global.App.Redis.Get(context.Background(), cacheKey).Val()
	} else {
		banner := models.Banner{}
		err := global.App.DB.Where("uid = ?", id).Order("id desc").First(&banner).Error
		if err != nil {
			return ""
		}
		mid := banner.MID
		media := models.Media{}
		err = global.App.DB.Where("id = ?", mid).Find(&media).Error
		if err != nil {
			return ""
		}
		url = "/storage" + "/" + media.Src
	}
	return url
}
