package controller

import (
	"blogBackend/models"
	"blogBackend/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		zap.L().Debug("c.ShouldBindJSON(post) err", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err := service.CreatePost(&post)
	if err != nil {
		zap.L().Error("service.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// PostListHandler 帖子分页展示
// 这里可以优化，与 FroPostListHandler 重复了；
func PostListHandler(c *gin.Context) {
	page, size := GetPageInfo(c)

	data, err := service.ShowPosts(page, size)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// FroPostListHandler 帖子分页展示
func FroPostListHandler(c *gin.Context) {
	page, size := GetPageInfo(c)

	data, err := service.FroShowPosts(page, size)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// ModifyPostHandler 修改帖子；
func ModifyPostHandler(c *gin.Context) {
	idStr := c.Param("id")
	var fo *models.ModifyForm
	err := c.ShouldBindJSON(&fo)
	if err != nil {
		zap.L().Error("modify post failed", zap.Error(err))
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err = service.ModifyPost(id, fo)
	if err != nil {
		zap.L().Error("service.modify post failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, nil)
}

// FindByIDHandler 此处遇到的问题：
// 2203 已解决：之前前端 axios 方法写错了，注释掉的三行为非必要；
func FindByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	//if len(idStr) == 0 {
	//	idStr = c.Param("id")
	//}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	data, err := service.ShowPostByID(id)
	if err != nil {
		zap.L().Error("service showPostByID failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, data)
}

// DelByIDHandler 通过 ID 删除；
func DelByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := service.DelPostByID(id)
	if err != nil {
		zap.L().Error("service DelByID failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, nil)
}
