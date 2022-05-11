package etcd

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/client/v3"
	"log"
	"logTransfer/conf"
	"logTransfer/es"
)

// Watcher 监控etcd中配置文件的变化；
func Watcher(key string) {
	for {
		watchCh := client.Watch(context.Background(), key)
		for wresp := range watchCh {
			for _, evt := range wresp.Events {
				log.Println("config changed")
				var newConfig []*conf.TransferEntry
				if evt.Type == clientv3.EventTypeDelete {
					// 如果是删除，则直接传一个空值；
					es.PutNewConf(newConfig)
					continue
				}
				err := json.Unmarshal(evt.Kv.Value, &newConfig)
				if err != nil {
					log.Println("watcher unmarshall json failed", err)
					continue
				}
				es.PutNewConf(newConfig)
			}
		}
	}
}
