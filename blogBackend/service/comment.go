package service

import (
	"blogBackend/dao/mysql"
	"blogBackend/models"
	"go.uber.org/zap"
)

// GetComment 获取评论数据；
func GetComment(id int64) (commentsData []*models.ApiResComments, err error) {
	commentsData = make([]*models.ApiResComments, 0, 2)
	comments, err := mysql.GetComment(id)
	if err != nil {
		zap.L().Error("service get comment failed", zap.Error(err))
		return
	}
	for _, comment := range comments {
		fromUserID := comment.FromUserID
		postID := comment.PostID
		// 这里是获取对于该评论的回复
		reply, _ := mysql.GetReply(int64(postID), int64(fromUserID))
		data := new(models.ApiResComments)
		data.Comment = comment
		data.ReplyContent = reply
		commentsData = append(commentsData, data)
	}
	return
}

// CreateComment 创建评论；
func CreateComment(id int64, idPrimary bool, comment *models.CreateCommentForm) (err error) {
	// 前端针对是否为一楼传入isPrimary;
	err = mysql.CreateComment(id, idPrimary, comment)
	if err != nil {
		zap.L().Error("service create comment failed", zap.Error(err))
		return
	}
	return
}
