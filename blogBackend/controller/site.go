package controller

import (
	"blogBackend/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetSiteInfo 获取网站信息；
func GetSiteInfo(c *gin.Context) {
	data, err := service.GetSiteInfo()
	if err != nil {
		zap.L().Error("controller get site info failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, data)
}
