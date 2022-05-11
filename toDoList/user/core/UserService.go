package core

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"user/model"
	"user/service"
)

// BuildUser 创建用户结构体实例；
func BuildUser(user model.User) *service.UserModel {
	userModel := service.UserModel{
		ID:       uint32(user.ID),
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
		UpdateAt: user.UpdatedAt.Unix(),
	}
	return &userModel
}

// UserLogin 用户登录；
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.Db.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	if user.CheckPassword(req.Password) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

// UserRegister 用户注册；
func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次输入密码不一致！")
		return err
	}
	count := 0
	if err := model.Db.Model(model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	user := model.User{
		UserName: req.UserName,
	}
	// 加密密码；
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := model.Db.Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
