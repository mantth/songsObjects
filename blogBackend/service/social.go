package service

import (
	"blogBackend/dao/mysql"
	"blogBackend/models"
	"go.uber.org/zap"
)

// GetSocialInfo 获取社交信息；
func GetSocialInfo() (site []*models.Social, err error) {
	site, err = mysql.GetSocialInfo()
	if err != nil {
		zap.L().Error("service get social info failed", zap.Error(err))
		return
	}
	return
}
