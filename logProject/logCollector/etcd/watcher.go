package etcd

import (
	"context"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"logProject/conf"
	"logProject/tails"
)

// Watcher 监控etcd中配置文件的变化；
func Watcher(key string) {
	// 这里要一直监听；
	for {
		watchCh := client.Watch(context.Background(), key)
		for wresp := range watchCh {
			for _, evt := range wresp.Events {
				log.Println("config changed")
				var newConfig []*conf.CollectEntry
				if evt.Type == clientv3.EventTypeDelete {
					// 如果是删除，则直接传一个空值；
					tails.PutNewConf(newConfig)
					continue
				}
				err := json.Unmarshal(evt.Kv.Value, &newConfig)
				if err != nil {
					log.Println("watcher unmarshall json failed", err)
					continue
				}
				// 将新配置传入配置文件chan;
				tails.PutNewConf(newConfig)
			}
		}
	}
}
