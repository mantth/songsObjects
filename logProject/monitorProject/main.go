package main

import (
	"monitorProject/influx"
	"monitorProject/sysMonitor"
)

// 未完成：kafka\etcd
func main() {
	influx.Init()
	var forever chan bool
	// 启用协程
	go sysMonitor.GetMemInfo()
	go sysMonitor.GetCPUInfo()
	go sysMonitor.GetDiskInfo()
	go influx.SendToInflux(false)
	// 阻塞main函数，持续收集和发送监控信息；
	<-forever
}
