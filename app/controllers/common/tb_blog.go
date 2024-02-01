package common

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/common/response"
	"bloggobackend/app/services"
	"github.com/gin-gonic/gin"
)

func BlogCreate(c *gin.Context) {
	var form request.Blog
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	data, _ := services.BlogService.CreateBlog(form)
	response.Success(c, data)
}

func BlogUpdate(c *gin.Context) {
	var form request.Blog
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	data, _ := services.BlogService.UpdateBlog(form)
	response.Success(c, data)
}

func BlogDelete(c *gin.Context) {
	var form request.Blog
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	data, _ := services.BlogService.DeleteBlog(form)
	response.Success(c, data)
}
