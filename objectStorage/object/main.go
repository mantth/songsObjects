package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/go-micro/plugins/v3/registry/etcd"
	"object/core"
	"object/service"
)

// 实际中要运行多个节点；
func main() {
	// rabbitMQ暂时弃用；
	//config := new(conf.Config)
	//configPath := "./conf/config.ini"
	//err := ini.MapTo(&config, configPath)
	//if err != nil {
	//	log.Printf("load config failed: %s ...", err)
	//	return
	//}
	//pathRabbitMQ := strings.Join([]string{config.Rabbit.RabbitMQ, "://", config.Rabbit.RabbitMQUser, ":",
	//	config.Rabbit.RabbitMQPassWord, "@", config.Rabbit.RabbitMQHost, ":", config.Rabbit.RabbitMQPort, "/"}, "")
	//core.Init(pathRabbitMQ)
	core.GetObjectInfo()
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	microService := micro.NewService(
		micro.Name("rpcFileService"),
		micro.Address("127.0.0.1:8081"),
		micro.Registry(etcdReg),
	)
	microService.Init()
	_ = service.RegisterFileServiceHandler(microService.Server(), new(core.FileService))
	_ = microService.Run()
}
