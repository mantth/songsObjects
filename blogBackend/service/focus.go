package service

import (
	"blogBackend/dao/mysql"
	"blogBackend/models"
	"go.uber.org/zap"
)

// GetFocus 获取热点；
func GetFocus() (focus []*models.Focus, err error) {
	focus, err = mysql.GetFocus()
	if err != nil {
		zap.L().Error("service get focus failed", zap.Error(err))
		return
	}
	return
}
