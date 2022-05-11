package core

import (
	"github.com/olivere/elastic/v7"
	"log"
)

var (
	cli *elastic.Client
	//msgChan chan interface{}
)

func Init(addr string) (err error) {
	opt := elastic.SetSniff(false)
	cli, err = elastic.NewClient(elastic.SetURL(addr), opt)
	if err != nil {
		return err
	}
	log.Println("init es client success...")
	//for _, index := range indexes {
	//	// 启用一个goroutine来发送消息到ES服务器‘
	//	in := index
	//	go func() {
	//		err := SendMsgToEs(in, ifLog)
	//		if err != nil {
	//			panic(err)
	//		}
	//	}()
	//}
	return
}
