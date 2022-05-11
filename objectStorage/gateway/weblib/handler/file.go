package handler

import (
	"context"
	"gateway/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// FileUpload 文件上传服务；
func FileUpload(c *gin.Context) {
	// 1. 调用fileService.isExist判断文件是否已存在；
	isExist, fileService := Locate(c)
	// 2. 往元数据服务中存入一条新的数据；
	PutMeta(c)
	// 如果不存在，写入，存在就直接返回200；
	if !isExist {
		var fileModel service.FileModel
		fileModel.FileName = c.Param("name")
		PanicIfFileError(c.Bind(&fileModel))
		var fileBinary service.FileBinary
		fileBinary.File = &fileModel
		fileBinary.File.HashCode = c.GetHeader("Digest")
		reader, _ := ioutil.ReadAll(c.Request.Body)
		fileBinary.Data = reader
		_, err := fileService.FileUpload(context.Background(), &fileBinary)
		PanicIfFileError(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// FileDownload 文件下载服务；
func FileDownload(c *gin.Context) {
	var fileRequest service.FileRequest
	// 这里因为实际保存的文件名是hash值，所以要先获取hash值
	meta := getObjectInfo(c)
	fileRequest.FileName = meta[0].Hash
	PanicIfFileError(c.Bind(&fileRequest))
	fileService := c.Keys["fileService"].(service.FileService)
	fileResp, err := fileService.FileDownload(context.Background(), &fileRequest)
	PanicIfFileError(err)
	fileBinary, err := fileResp.Recv()
	PanicIfFileError(err)
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"msg": string(fileBinary.Data),
		},
	})
}
