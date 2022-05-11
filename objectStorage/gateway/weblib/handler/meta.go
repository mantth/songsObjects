package handler

import (
	"context"
	"gateway/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetMeta 获取元数据；
func GetMeta(c *gin.Context) {
	var metaReq service.MetaRequest
	PanicIfMetaError(c.Bind(&metaReq))
	metaReq.Name = c.Param("name")
	versionStr := c.Query("version")
	if len(versionStr) != 0 {
		version, err := strconv.ParseInt(versionStr, 10, 64)
		PanicIfMetaError(err)
		metaReq.Version = int32(version)
	} else {
		metaReq.Version = -2
	}
	metaService := c.Keys["metaService"].(service.MetaService)
	meta, err := metaService.GetVersion(context.Background(), &metaReq)
	PanicIfMetaError(err)
	c.JSON(http.StatusOK, gin.H{
		"meta": meta.Meta,
	})
}

// PutMeta 增加元数据；
func PutMeta(c *gin.Context) {
	var metaReq service.MetaRequest
	PanicIfMetaError(c.Bind(&metaReq))
	metaReq.Name = c.Param("name")
	metaReq.Hash = c.GetHeader("Digest")
	length, err := strconv.ParseUint(c.GetHeader("Content-Length"), 10, 64)
	PanicIfMetaError(err)
	metaReq.Length = length
	metaService := c.Keys["metaService"].(service.MetaService)
	_, err = metaService.PutVersion(context.Background(), &metaReq)
	PanicIfMetaError(err)
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "ok",
	//})
}

// DelMeta 删除元数据；
// 此处并不是真正的删除，只是将version、hash置零;
func DelMeta(c *gin.Context) {
	var metaReq service.MetaRequest
	PanicIfMetaError(c.Bind(&metaReq))
	metaReq.Name = c.Param("name")
	metaService := c.Keys["metaService"].(service.MetaService)
	_, err := metaService.PutVersion(context.Background(), &metaReq)
	PanicIfMetaError(err)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
