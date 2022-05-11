package models

import (
	"encoding/json"
	"errors"
	"time"
)

// Post Post 类型；
type Post struct {
	ID            uint64    `json:"id,string" db:"id"`
	Banner        string    `json:"banner" db:"banner"`
	IsTop         int       `json:"isTop,string" db:"is_top"`
	IsHot         int       `json:"isHot,string" db:"is_hot"`
	PubTime       time.Time `json:"pubTime" db:"pub_time"`
	Title         string    `json:"title" db:"title" binding:"required"`
	Summary       string    `json:"summary" db:"summary" binding:"required"`
	Content       string    `json:"content" db:"content" binding:"required"`
	ViewsCount    uint64    `json:"viewsCount,string" db:"views_count"`
	CommentsCount uint64    `json:"commentsCount,string" db:"comments_count"`
	Type          string    `json:"type" db:"type" binding:"required"`
}

// ModifyForm 前端修改帖子传入的表单数据；
type ModifyForm struct {
	Banner  string    `json:"banner" db:"banner"`
	IsTop   int       `json:"isTop,string" db:"is_top"`
	IsHot   int       `json:"isHot,string" db:"is_hot"`
	PubTime time.Time `json:"pubTime" db:"pub_time"`
	Title   string    `json:"title" db:"title" binding:"required"`
	Summary string    `json:"summary" db:"summary" binding:"required"`
	Content string    `json:"content" db:"content" binding:"required"`
	Type    string    `json:"type" db:"type" binding:"required"`
}

// UnmarshalJson 为Post类型实现UnmarshalJson方法；
func (p *Post) UnmarshalJson(data []byte) (err error) {
	required := struct {
		Title   string `json:"title" db:"title"`
		Summary string `json:"summary" db:"summary"`
		Content string `json:"content" db:"content"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("帖子标题不能为空")
	} else if len(required.Content) == 0 {
		err = errors.New("帖子内容不能为空")
	} else if len(required.Summary) == 0 {
		err = errors.New("未输入文章概要")
	} else {
		p.Title = required.Title
		p.Summary = required.Summary
		p.Content = required.Content
	}
	return
}

func (fo *ModifyForm) UnmarshalJson(data []byte) (err error) {
	required := struct {
		Banner  string `json:"banner" db:"banner"`
		IsTop   int    `json:"isTop,string" db:"is_top"`
		IsHot   int    `json:"isHot,string" db:"is_hot"`
		Title   string `json:"title" db:"title"`
		Summary string `json:"summary" db:"summary"`
		Content string `json:"content" db:"content"`
		Type    string `json:"type" db:"type"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("帖子标题不能为空")
	} else if len(required.Content) == 0 {
		err = errors.New("帖子内容不能为空")
	} else if len(required.Summary) == 0 {
		err = errors.New("未输入文章概要")
	} else {
		fo.IsTop = required.IsTop
		fo.IsHot = required.IsHot
		fo.Title = required.Title
		fo.Summary = required.Summary
		fo.Content = required.Content
		fo.Type = required.Type
	}
	return
}

// ResPostsList 帖子分页数据之封装；
type ResPostsList struct {
	Total       int64   `json:"total"`
	Posts       []*Post `json:"posts"`
	Page        int64   `json:"page"`
	HasNextPage bool    `json:"hasNextPage"`
}
