package etcd

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/client/v3"
	"log"
	"logTransfer/conf"
	"time"
)

var (
	client *clientv3.Client
)

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

func GetConf(key string) (transEntryList []*conf.TransferEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		log.Printf("get config from etcd failed: %s", err)
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		log.Println("get nothing by the key")
		return nil, err
	}
	ret := resp.Kvs[0]
	err = json.Unmarshal(ret.Value, &transEntryList)
	if err != nil {
		log.Println("config unmarshall err")
		return nil, err
	}
	return
}
