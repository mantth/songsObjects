package main

import (
	"gateway/service"
	"gateway/weblib"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/web"
	"github.com/go-micro/plugins/v3/registry/etcd"
	"time"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	fileMicroService := micro.NewService(
		micro.Name("fileService.Client"),
	)

	metaMicroService := micro.NewService(
		micro.Name("metaService.Client"),
	)
	fileService := service.NewFileService("rpcFileService", fileMicroService.Client())
	metaService := service.NewMetaService("rpcMetaService", metaMicroService.Client())

	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		web.Handler(weblib.NewRouter(fileService, metaService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	_ = server.Init()
	_ = server.Run()
}
