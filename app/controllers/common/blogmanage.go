package common

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/common/response"
	"bloggobackend/app/services"
	"github.com/gin-gonic/gin"
)

func BlogManage(c *gin.Context) {
	var form request.BlogManage
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	data := services.BlogManageService.GetBlogByPage(form)
	c.JSON(200, data)
}
