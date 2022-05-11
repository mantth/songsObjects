package handler

import (
	"context"
	"gateway/service"
	"github.com/gin-gonic/gin"
	"github.com/klauspost/reedsolomon"
	"io"
)

// rs encoder
type encoder struct {
	readers []io.Reader
	enc     reedsolomon.Encoder
	cache   []byte
}

// Locate 调用fileService微服务，判断文件是否存在
func Locate(c *gin.Context) (bool, service.FileService) {
	var fileReq service.FileRequest
	fileReq.FileName = c.GetHeader("Digest")
	PanicIfFileError(c.Bind(&fileReq))
	fileService := c.Keys["fileService"].(service.FileService)
	fileResp, err := fileService.FileExist(context.Background(), &fileReq)
	PanicIfFileError(err)
	return fileResp.File.IsExist, fileService
}

// 获取对象信息；
//ToDO: 此处有bug，version 传入应校验;
func getObjectInfo(c *gin.Context) []*service.MetaData {
	var metaReq service.MetaRequest
	PanicIfMetaError(c.Bind(&metaReq))
	metaReq.Name = c.Param("name")
	metaReq.Version = -2
	metaService := c.Keys["metaService"].(service.MetaService)
	meta, err := metaService.GetVersion(context.Background(), &metaReq)
	PanicIfMetaError(err)
	return meta.Meta
}

// ToDO: 创建数据分片，遍历写入各个节点;
//func PutShard(dataServers []string, hash string, data []byte) error {
//	if len(dataServers) != 6 {
//		return errors.New("dataServers number mismatch")
//	}
//	perShard := (len(data) + 4 - 1) / 4
//	readers := make([]io.Reader, 6)
//	for i := range readers {
//		readers[i] = bytes.NewReader(data)
//	}
//
//}
