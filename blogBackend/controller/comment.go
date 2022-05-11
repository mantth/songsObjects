package controller

import (
	"blogBackend/models"
	"blogBackend/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetComment 根据帖子 ID 获取其下所有评论；
func GetComment(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	comments, err := service.GetComment(id)
	if err != nil {
		zap.L().Error("controller get comment failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, comments)
}

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	var comment models.CreateCommentForm
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	primaryStr := c.Query("isPrimary")
	//fmt.Println(idStr)
	isPrimary, _ := strconv.ParseBool(primaryStr)
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		zap.L().Debug("c.ShouldBindJSON(&comment) err", zap.Any("err", err))
		zap.L().Error("create comment with invalid param")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err = service.CreateComment(id, isPrimary, &comment)
	if err != nil {
		zap.L().Error("service.CreateComment failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
