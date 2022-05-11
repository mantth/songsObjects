package models

import "time"

// Comment 评论结构体；
type Comment struct {
	UserAvatar   string    `json:"fromUserAvatar" db:"user_avatar"`
	ID           uint64    `json:"id" db:"id"`
	PostID       uint64    `json:"postID" db:"post_id"`
	ParentID     uint64    `json:"parentID" db:"parent_id"`
	FromUserID   uint64    `json:"fromUserID" db:"from_user_id"`
	FromUserName string    `json:"fromUserName" db:"from_username"`
	Content      string    `json:"content" db:"content"`
	CreateTime   time.Time `json:"createTime" db:"create_time"`
	ToUserID     uint64    `json:"toUserID" db:"to_user_id"`
	ToUserName   string    `json:"toUserName" db:"to_username"`
}

// CreateCommentForm 前端表单数据；
type CreateCommentForm struct {
	PostID       uint64 `json:"postID" db:"post_id"`
	FromUserID   uint64 `json:"fromUserID" db:"from_user_id"`
	FromUserName string `json:"fromUserName" db:"from_username"`
	Content      string `json:"content" db:"content"`
	ToUserID     uint64 `json:"toUserID" db:"to_user_id"`
	ToUserName   string `json:"toUserName" db:"to_username"`
}

// ApiResComments 封装评论请求返回数据
type ApiResComments struct {
	Comment      *Comment   `json:"comment"`
	ReplyContent []*Comment `json:"reply"`
}
