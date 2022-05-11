package mysql

import (
	"blogBackend/models"
	"go.uber.org/zap"
)

// GetSocialInfo 获取 QQ、github 等信息；
func GetSocialInfo() (site []*models.Social, err error) {
	socials := make([]*models.Social, 0, 2)
	sqlStr := `select * from social`
	err = db.Select(&socials, sqlStr)
	if err != nil {
		zap.L().Error("mysql get social info failed", zap.Error(err))
		return
	}
	return
}
