package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/go-ini/ini"
	"github.com/go-micro/plugins/v3/registry/etcd"
	"log"
	"metadata/conf"
	"metadata/core"
	"metadata/service"
)

func main() {
	config := new(conf.Config)
	configPath := "./conf/config.ini"
	err := ini.MapTo(&config, configPath)
	if err != nil {
		log.Printf("load config failed: %s ...", err)
		return
	}
	err = core.Init(config.EsConf.Address)
	if err != nil {
		log.Println(err)
		return
	}
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	microService := micro.NewService(
		micro.Name("rpcMetaService"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg),
	)
	microService.Init()
	_ = service.RegisterMetaServiceHandler(microService.Server(), new(core.MetaService))
	_ = microService.Run()
}
