package main

import (
	"github.com/go-ini/ini"
	"log"
	"logTransfer/conf"
	"logTransfer/es"
	"logTransfer/etcd"
	"logTransfer/kafka"
)

func main() {
	var forever chan bool
	config := new(conf.Config)
	configPath := "./conf/config.ini"
	err := ini.MapTo(&config, configPath)
	if err != nil {
		log.Printf("load config failed: %s ...", err)
		return
	}
	err = kafka.Init(config.KafkaConf.Address)
	if err != nil {
		log.Printf("init kafka failed: %s ...", err)
		return
	}
	log.Println("init kafka success...")
	err = es.InitConn(config.EsConf.Address)

	if err != nil {
		log.Printf("init es connection failed: %s ...", err)
		return
	}
	err = etcd.Init(config.EtcdConf.Address)
	if err != nil {
		log.Printf("init etcd failed: %s", err)
		return
	}
	log.Println("init etcd success...")
	entryList, err := etcd.GetConf(config.EtcdConf.Key)
	if err != nil {
		log.Printf("get config from etcd failed: %s", err)
		return
	}
	go etcd.Watcher(config.EtcdConf.Key)
	es.InitMsgChan(entryList, config.EsConf.Size)
	es.Init(entryList, config.EsConf.Size)
	log.Println("init elasticsearch success...")
	<-forever
}
