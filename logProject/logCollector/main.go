package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"logProject/conf"
	"logProject/etcd"
	"logProject/kafka"
	"logProject/service"
	"logProject/tails"
)

func main() {
	// 获取本机IP;
	ip, err := conf.GetLocalIP()
	if err != nil {
		log.Printf("get local ip failed: %s", err)
		return
	}
	// 初始化配置类；
	config := new(conf.Config)
	configPath := "./conf/config.ini"
	err = ini.MapTo(&config, configPath)
	if err != nil {
		log.Printf("load config failed: %s", err)
		return
	}
	err = kafka.Init([]string{config.KafakConfig.Address}, config.KafakConfig.ChanSize)
	if err != nil {
		log.Printf("init kafka failed: %s", err)
		return
	}
	log.Println("init kafka success!")
	err = etcd.Init(config.EtcdConfig.Address)
	if err != nil {
		log.Printf("init etcd failed: %s", err)
		return
	}
	// 将本机IP与etcd logPath 拼接得到实际的logPath;
	logPath := fmt.Sprintf(config.EtcdConfig.LogPath, ip)
	entryList, err := etcd.GetConf(logPath)
	if err != nil {
		log.Printf("get config from etcd failed: %s", err)
		return
	}
	// 监听配置文件变化；
	go etcd.Watcher(logPath)
	//fmt.Println(entryList[1].Topic)
	err = tails.Init(entryList)
	if err != nil {
		log.Printf("init tail failed: %s", err)
		return
	}
	log.Println("init tail success!")
	service.Run()
}
