package mysql

import (
	"blogBackend/models"
	"go.uber.org/zap"
)

// GetSiteInfo 获取网站信息；
func GetSiteInfo() (site *models.Site, err error) {
	sites := make([]*models.Site, 0, 2)
	sqlStr := `select * from site`
	err = db.Select(&sites, sqlStr)
	if err != nil {
		zap.L().Error("mysql get site info failed", zap.Error(err))
		return
	}
	return sites[0], err
}
