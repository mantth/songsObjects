package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/service"
	"context"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册；
func UserRegister(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))
	userService := c.Keys["userService"].(service.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	c.JSON(200, gin.H{
		"data": userResp,
	})
}

// UserLogin 用户登录；
func UserLogin(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))
	userService := c.Keys["userService"].(service.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	c.JSON(200, gin.H{
		"code": userResp.Code,
		"msg":  "success",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}
