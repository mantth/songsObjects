package main

import (
	"rabbit_server/conf"
	"rabbit_server/service"
)

func main() {
	conf.Init()
	forever := make(chan bool)
	service.CreateTask()
	// 阻塞
	<-forever
}
