package sysMonitor

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"monitorProject/influx"
	"time"
)

// GetCPUInfo 获取CPU监控信息；
func GetCPUInfo() {
	// 阻塞，两秒钟获取一次；
	for {
		cpuInfo, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Printf("get cpu info failed: %s", err)
			return
		}
		// 初始化cpu_info类型的一个消息；
		newCPUInfo := &influx.NewPoint{}
		newCPUInfo.Measurement = "cpu_info"
		newCPUInfo.Tags = map[string]string{"unit": "percent"}
		newCPUInfo.Fields = map[string]interface{}{"cpu_usage": cpuInfo[0]}
		//fmt.Println(cpuInfo[0])
		// 发送到消息管道中；
		influx.SendInfoToChan(newCPUInfo)
		time.Sleep(time.Second * 2)
	}
}

func GetMemInfo() {
	for {
		memInfo, err := mem.VirtualMemory()
		if err != nil {
			log.Printf("get cpu info failed: %s", err)
			return
		}
		newMemInfo := &influx.NewPoint{}
		newMemInfo.Measurement = "mem_info"
		newMemInfo.Tags = map[string]string{"unit": "percent"}
		newMemInfo.Fields = map[string]interface{}{"mem_usage": memInfo.UsedPercent}
		influx.SendInfoToChan(newMemInfo)
		time.Sleep(time.Second * 2)
	}
}

// GetDiskInfo 获取磁盘信息；
func GetDiskInfo() {
	// 阻塞，每20秒获取一次；
	for {
		diskInfo, err := disk.Partitions(false)
		if err != nil {
			log.Printf("get disk partions info failed: %s", err)
			return
		}
		// 遍历各个磁盘的信息；
		for _, part := range diskInfo {
			diskUsage, err := disk.Usage(part.Mountpoint)
			if err != nil {
				log.Printf("get disk info failed: %s", err)
				continue
			}
			//fmt.Println(part.Mountpoint)
			newDIskInfo := &influx.NewPoint{}
			newDIskInfo.Measurement = "disk_info"
			newDIskInfo.Tags = map[string]string{"disk": "disk_usage"}
			newDIskInfo.Fields = map[string]interface{}{part.Mountpoint: diskUsage.UsedPercent}
			influx.SendInfoToChan(newDIskInfo)
		}
		time.Sleep(time.Second * 10)
	}
}
