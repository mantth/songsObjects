package service

import (
	"blogBackend/dao/mysql"
	"blogBackend/models"
	"go.uber.org/zap"
)

// GetSiteInfo 获取网站信息；
func GetSiteInfo() (site *models.Site, err error) {
	site, err = mysql.GetSiteInfo()
	if err != nil {
		zap.L().Error("service get sit info failed", zap.Error(err))
		return
	}
	return
}
