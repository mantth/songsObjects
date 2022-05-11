package etcd

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/client/v3"
	"log"
	"logProject/conf"
	"time"
)

var (
	client *clientv3.Client
)

// Init 初始化etcd客户端；
func Init(address string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{address},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Printf("init etcd conn failed: %s", err)
		return err
	}
	return
}

// GetConf 从etcd中取得配置信息；
func GetConf(key string) (collectEntryList []*conf.CollectEntry, err error) {
	// 设定一个两秒的等待时间；
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		log.Printf("get config from etcd failed: %s", err)
		return nil, err
	}
	// 没获取到就返回空；
	if len(resp.Kvs) == 0 {
		log.Println("get nothing by the key")
		return nil, err
	}
	// 这里只需要kv中的第一项（只有一项）；
	ret := resp.Kvs[0]
	err = json.Unmarshal(ret.Value, &collectEntryList)
	if err != nil {
		log.Println("config unmarshall err")
		return nil, err
	}
	return
}
