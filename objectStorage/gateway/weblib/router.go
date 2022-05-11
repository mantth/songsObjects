package weblib

import (
	"gateway/weblib/handler"
	"gateway/weblib/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middlewares.Cors(), middlewares.InitMiddleware(service), middlewares.ErrorMiddleware())
	v1 := ginRouter.Group("api/v1")
	{
		v1.GET("object/:name", handler.FileDownload)
		v1.PUT("object/:name", handler.FileUpload)
		v1.GET("version/:name", handler.GetMeta)
		v1.PUT("version/:name", handler.PutMeta)
		v1.DELETE("version/:name", handler.DelMeta)
	}
	return ginRouter

}
