package common

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/common/response"
	"bloggobackend/app/services"
	"github.com/gin-gonic/gin"
)

func BannerGet(c *gin.Context) {
	var form request.BannerGet
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	url := services.BannerService.GetUrlByID(form)
	response.Success(c, url)
}
