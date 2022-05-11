package mysql

import (
	"blogBackend/models"
	"go.uber.org/zap"
)

func GetReply(id, toUserID int64) (reply []*models.Comment, err error) {
	sqlStr := "select * from comment where post_id = (?) and to_user_id =(?)"
	reply = make([]*models.Comment, 0, 2)
	err = db.Select(&reply, sqlStr, id, toUserID)
	if err != nil {
		zap.L().Error("mysql get reply failed", zap.Error(err))
		return
	}
	return
}
