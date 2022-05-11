package influx

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"log"
	"time"
)

// 初始化全局变量；
var (
	bucket   string
	org      string
	token    string
	url      string
	cli      influxdb2.Client
	infoChan chan *NewPoint
)

// Init 初始化influxDB连接和信息通道；
func Init() {
	bucket = "logBucket"
	org = "root"
	token = "Upvr2vjTPzErdKvcmsiH6xzlSWWUHwZ3OkO1bg71a35uI7KWjKJ8oCo9K1yeEZE-DwaHc6wcJR8sgHdeCf-NLw=="
	url = "http://localhost:8086"
	cli = influxdb2.NewClient(url, token)
	infoChan = make(chan *NewPoint, 100)
	log.Println("init influxDB conn success")
	defer cli.Close()
}

// SendToInflux 将监控消息发送到influxDB数据库中；
func SendToInflux(ifLog bool) {
	writerAPI := cli.WriteAPIBlocking(org, bucket)
	log.Println("start to send info to influxDB...")
	for {
		select {
		case newInfo := <-infoChan:
			p := influxdb2.NewPoint(
				newInfo.Measurement,
				newInfo.Tags,
				newInfo.Fields,
				time.Now(),
			)
			err := writerAPI.WritePoint(context.Background(), p)
			if err != nil {
				log.Println("send to influxDB failed: ", err, newInfo.Measurement)
				return
			}
			if ifLog {
				log.Printf("send to influxDB success: %s \n", newInfo.Measurement)
			}

		}
	}
}

// SendInfoToChan 为避免暴露infoChan，为monitor开放一个将消息存入channel的方法；
func SendInfoToChan(point *NewPoint) {
	infoChan <- point
}
