package core

import (
	"context"
	"log"
	"object/service"
)

// FileUpload 文件上传服务；
func (*FileService) FileUpload(ctx context.Context, binary *service.FileBinary, response *service.FileResponse) error {
	err := checkAndSaveFile(binary)
	if err != nil {
		log.Println("file service save file failed: ", err)
		return err
	}
	return nil
}

// FileDownload 文件下载服务；
func (*FileService) FileDownload(ctx context.Context, request *service.FileRequest, stream service.FileService_FileDownloadStream) error {
	fileByte, err := getFile(request)
	if err != nil {
		log.Println("get file failed: ", err)
		return err
	}
	fileBinary := service.FileBinary{
		Data: fileByte,
	}
	err = stream.Send(&fileBinary)
	if err != nil {
		log.Println("send file failed:", err)
		return err
	}
	return nil
}

// FileExist 判断文件是否存在
// 此处locate会扫描目录，将文件信息存储在map中以避免频繁的磁盘读写操作；
func (*FileService) FileExist(ctx context.Context, request *service.FileRequest, response *service.FileResponse) error {
	stat := Locate(request.HashCode)
	var fileModel service.FileModel
	fileModel.IsExist = stat
	response.File = &fileModel
	return nil
}
