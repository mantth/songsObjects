package controller

import (
	"blogBackend/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetSocialInfo 获取社交信息；
func GetSocialInfo(c *gin.Context) {
	data, err := service.GetSocialInfo()
	if err != nil {
		zap.L().Error("controller get social info failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, data)
}
