package mysql

import (
	"blogBackend/models"
	"go.uber.org/zap"
	"strconv"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post(is_top, is_hot, banner, title, summary, content, type) values(?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, post.IsTop, post.IsHot, post.Banner, post.Title, post.Summary, post.Content, post.Type)
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

// GetPostList 获取帖子数据
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select id, banner, is_top, is_hot, pub_time, title, summary, content, type
	from post ORDER BY pub_time DESC limit ?,?
	`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	if err != nil {
		zap.L().Error("get post list failed", zap.Error(err))
		return
	}
	return
}

// GetTotalNum 获取帖子数量；
func GetTotalNum() (data int64) {
	sqlStr := `SELECT COUNT(*) FROM post`
	total := []int64{}
	err := db.Select(&total, sqlStr)
	if err != nil {
		zap.L().Error("get total failed", zap.Error(err))
		return
	}
	return total[0]
}

// FindByIDAndUpdate 通过ID更新帖子
func FindByIDAndUpdate(id int64, fo *models.ModifyForm) (err error) {
	sqlStr := `update post set banner=(?), is_hot=(?), is_top=(?), title=(?), summary=(?), content=(?), type=(?) where id = (?)`
	_, err = db.Exec(sqlStr, fo.Banner, fo.IsHot, fo.IsTop, fo.Title, fo.Summary, fo.Content, fo.Type, id)
	if err != nil {
		return
	}
	return
}

// FindByIDAndDel 通过ID删除帖子
func FindByIDAndDel(id int64) (err error) {
	sqlStr := `delete from post where id = (?)`
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		zap.L().Error("get post by Id failed", zap.Error(err))
		return
	}
	return
}

// FindByID 通过ID查找帖子
func FindByID(id int64) (modifyPost []*models.ModifyForm, err error) {
	modifyPost = make([]*models.ModifyForm, 0, 2)
	sqlStr := `select banner, is_top, is_hot, pub_time, title, summary, content, type
	from post where id = (?)`
	err = db.Select(&modifyPost, sqlStr, strconv.FormatInt(id, 10))
	//fmt.Println(modifyPost)
	if err != nil {
		zap.L().Error("get post by Id failed", zap.Error(err))
		return
	}
	return
}
