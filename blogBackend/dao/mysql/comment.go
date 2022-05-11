package mysql

import (
	"blogBackend/models"
	"go.uber.org/zap"
)

// GetComment 根据传入的帖子ID获取所有评论；
func GetComment(id int64) (comments []*models.Comment, err error) {
	sqlStr := `select * from comment where post_id = (?) and to_user_id = 0`
	comments = make([]*models.Comment, 0, 2)
	err = db.Select(&comments, sqlStr, id)
	if err != nil {
		zap.L().Error("mysql get comment failed", zap.Error(err))
		return
	}
	return
}

// GetCountByPost 根据传入的帖子ID计算该帖子的评论数；
func GetCountByPost(id int64) (count int, err error) {
	total := []int{}
	//fmt.Println(id)
	sqlStr := `SELECT COUNT(content) from comment where post_id=(?)`
	err = db.Select(&total, sqlStr, id)
	if err != nil {
		zap.L().Error("mysql get comment count failed", zap.Error(err))
		return 0, err
	}
	return total[0], err
}

// CreateComment 此处设计为每个帖子下的评论根据先后生成FromUserID(1, 2, 3...0为保留字，代表该评论为主评论)；
// isPrimary 用来判断该评论属于回复还是主评论；
func CreateComment(id int64, isPrimary bool, comment *models.CreateCommentForm) (err error) {
	var sqlStr string
	UserID, _ := GetCountByPost(id)
	// 为主评论时， ToUserID默认为0；
	if isPrimary {
		sqlStr = `insert into comment (post_id, from_user_id, from_username, content) values(?, ?, ?, ?)`
		_, err = db.Exec(sqlStr, comment.PostID, UserID+1, comment.FromUserName, comment.Content)
		return
	}
	// 不为主评论时， 则传入对应的ToUserID， reply comment 时执行；
	sqlStr = `insert into comment (post_id, from_user_id, content, to_user_id)
	values(?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, id, UserID+1, comment.Content, comment.ToUserID)
	if err != nil {
		zap.L().Error("mysql create comment failed", zap.Error(err))
		return
	}
	return
}
