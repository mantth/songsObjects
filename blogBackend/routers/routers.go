package routers

import (
	"blogBackend/controller"
	"blogBackend/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter 建立路由；
func SetupRouter(mode string) *gin.Engine {
	// 设置成发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// 使用日志记录中间件和错误处理中间件；
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	v1.POST("/article/create", controller.CreatePostHandler)
	v1.GET("/article/index", controller.PostListHandler)
	v1.PUT("/article/:id", controller.ModifyPostHandler)
	v1.GET("/article/:id", controller.FindByIDHandler)
	v1.DELETE("/article/:id", controller.DelByIDHandler)
	v1.GET("post/index", controller.FroPostListHandler)
	v1.GET("/focus/list", controller.FocusHandler)
	v1.GET("/site", controller.GetSiteInfo)
	v1.GET("/social", controller.GetSocialInfo)
	v1.GET("/comment/:id", controller.GetComment)
	v1.PUT("/comment/:id", controller.CreateComment)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
