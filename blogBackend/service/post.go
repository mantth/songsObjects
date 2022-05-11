package service

import (
	"blogBackend/dao/mysql"
	"blogBackend/models"
	"fmt"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	if err := mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}
	return
}

// ShowPosts 后台用帖子列表
// 220211补：这里要优化，不必分前后台
func ShowPosts(page, size int64) (posts []*models.Post, err error) {
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, post := range postList {
		posts = append(posts, post)
	}
	return posts, nil
}

// FroShowPosts 前台帖子列表
func FroShowPosts(page, size int64) (data *models.ResPostsList, err error) {
	postList, err := mysql.GetPostList(page, size)
	total := mysql.GetTotalNum()
	if err != nil {
		fmt.Println(err)
		return
	}
	data = new(models.ResPostsList)
	data.Posts = postList
	data.Page = page
	data.Total = total
	// 前端代码已经修改，此处 hasNextPage 已无实际作用，可以去除；
	if int64(len(postList)) < size || postList == nil {
		data.HasNextPage = false
	}
	data.HasNextPage = true
	return data, nil
}

// ModifyPost 修改帖子；
func ModifyPost(id int64, fo *models.ModifyForm) (err error) {
	if err := mysql.FindByIDAndUpdate(id, fo); err != nil {
		zap.L().Error("mysql.FindByIdAndUpdate failed", zap.Error(err))
		return err
	}
	return
}

// ShowPostByID 根据 ID 返回帖子数据；
func ShowPostByID(id int64) (fo []*models.ModifyForm, err error) {
	fo, err = mysql.FindByID(id)
	if err != nil {
		zap.L().Error("mysql.FindById failed", zap.Error(err))
		return
	}
	return
}

// DelPostByID 通过 ID 删除；
func DelPostByID(id int64) (err error) {
	if err := mysql.FindByIDAndDel(id); err != nil {
		zap.L().Error("mysql.FindByIdAndDel failed", zap.Error(err))
		return err
	}
	return
}
