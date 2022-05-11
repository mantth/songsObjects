package mysql

import (
	"blogBackend/models"
	"go.uber.org/zap"
)

// GetFocus 获取热点；
func GetFocus() (focus []*models.Focus, err error) {
	sqlStr := "select * from focus"
	focus = make([]*models.Focus, 0, 2)
	err = db.Select(&focus, sqlStr)
	if err != nil {
		zap.L().Error("mysql get focus failed", zap.Error(err))
		return
	}
	return
}
