package models

import (
	"encoding/json"
	"errors"
)

type Focus struct {
	ID    uint64 `json:"id,string" db:"id"`
	Title string `json:"title" db:"title"`
	Img   string `json:"img" db:"img"`
}

// UnmarshalJson 为Focus类型实现UnmarshalJson方法；
func (focus *Focus) UnmarshalJson(data []byte) (err error) {
	required := struct {
		Title string `json:"title" db:"title"`
		Img   string `json:"img" db:"img"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("标题不能为空")
	} else if len(required.Img) == 0 {
		err = errors.New("内容不能为空")
	} else {
		focus.Title = required.Title
		focus.Img = required.Img
	}
	return
}
