package core

import (
	"context"
	"log"
	"metadata/service"
)

func (*MetaService) GetVersion(ctx context.Context, request *service.MetaRequest, response *service.MetaResponse) error {
	result := getMetaFromEs(request.Name, request.Version)
	err, data := getResult(result)
	if err != nil {
		return err
	}
	response.Meta = data
	return nil
}

func (*MetaService) DelVersion(ctx context.Context, request *service.MetaRequest, response *service.MetaResponse) error {
	metas := getLatestVersion(request.Name)
	err := sendToEs(request.Name, "", 0, metas[0].Version+1)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (*MetaService) PutVersion(ctx context.Context, request *service.MetaRequest, response *service.MetaResponse) error {
	metas := getLatestVersion(request.Name)
	err := sendToEs(request.Name, request.Hash, request.Length, metas[0].Version)
	//fmt.Println(request.Name)
	if err != nil {
		return err
	}
	return nil
}

func (*MetaService) GetLatestVersion(ctx context.Context, request *service.MetaRequest, response *service.MetaResponse) error {
	metas := getLatestVersion(request.Name)
	response.Meta = metas
	return nil
}
