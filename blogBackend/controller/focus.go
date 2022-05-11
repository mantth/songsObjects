package controller

import (
	"blogBackend/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// FocusHandler 获取热点轮播；
func FocusHandler(c *gin.Context) {
	focus, err := service.GetFocus()
	if err != nil {
		zap.L().Error("controller FocusHandler failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, focus)
}
