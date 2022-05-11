package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

// NewRouter 建立路由；
func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	// 使用中间件；
	ginRouter.Use(middlewares.Cors(), middlewares.InitMiddleware(service), middlewares.ErrorMiddleware())
	// session相关
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mySession", store))
	v1 := ginRouter.Group("api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		v1.POST("user/register", handlers.UserRegister)
		v1.POST("user/login", handlers.UserLogin)
		// 以下操作都需要用户验证；
		authed := v1.Group("/")
		authed.Use(middlewares.JWT())
		{
			authed.GET("tasks", handlers.GetTaskList)
			authed.POST("task", handlers.CreateTask)
			authed.GET("task/:id", handlers.GetTaskDetail) // task_id
			authed.PUT("task/:id", handlers.UpdateTask)    // task_id
			authed.DELETE("task/:id", handlers.DeleteTask) // task_id
		}
	}
	return ginRouter
}
