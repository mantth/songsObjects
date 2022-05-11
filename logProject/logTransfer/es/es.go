package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
	"logTransfer/conf"
)

// 全局客户端；
var (
	cli *elastic.Client
)

/**
此处针对每个topic创建一个消息通道并保存在map中方便控制；
key: topic;
val: channel;
*/
type msgChanMap struct {
	// map中存放指针类型，节省内存；
	chMap map[string]*chan interface{}
}

// 定义一个全局chanMap便于管理；
var (
	chanMap *msgChanMap
)

// elasticsearch任务结构体，封装context以能够控制协程中断；
type esTask struct {
	index  string
	topic  string
	ctx    context.Context
	cancel context.CancelFunc
}

// Run esTask调用之后，将从对应 topic 的 channel 中取数据发送到 es 服务中的对应 index 中，
// 该 index 即为 es.topic;
// ifLog 用于控制是否在控制台输出日志；
func (es *esTask) Run(ifLog bool) {
	for {
		select {
		case <-es.ctx.Done():
			log.Printf("esTask %s has stopped...", es.topic)
			return
		case msg := <-*chanMap.chMap[es.topic]:
			// 此处根据消息对应的topic发送到不同的index中；
			response, err := cli.Index().Index(es.topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				log.Println(err)
				return
			}
			if ifLog {
				//fmt.Println(msg)
				log.Printf("ID %s to index %s, type %s \n", response.Id, response.Index, response.Type)
			}
		}

	}
}

// InitConn 初始化elasticsearch连接；
func InitConn(addr string) (err error) {
	// 初始化消息通道；
	chanMap = new(msgChanMap)
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

// InitMsgChan 初始化消息通道；
func InitMsgChan(conf []*conf.TransferEntry, size int) {
	tempMap := make(map[string]*chan interface{})
	for _, config := range conf {
		newChan := make(chan interface{}, size)
		// 使用 topic 作为 key;
		tempMap[config.Topic] = &newChan
	}
	chanMap.chMap = tempMap
}

// SendMsgToChan 开放给 kafka 模块使用，将从 kafka 中取得的消息存入对应的 channel 中；
func SendMsgToChan(msg interface{}, topic string) {
	*chanMap.chMap[topic] <- msg
}

// newMsgChan 根据配置文件中 topic 新建消息通道；
// 此处 size 可优化，从配置文件中获取；
func newMsgChan(topic string) {
	newMsgCh := make(chan interface{}, 1000)
	chanMap.chMap[topic] = &newMsgCh
}

// ChanExist 判断该 topic 对应的消息通道是否存在；
// 此处是为了防止 kafka 给未创建的通道发消息；
func ChanExist(topic string) bool {
	_, ok := chanMap.chMap[topic]
	return ok
}

func closeMsgChan(topic string) {
	close(*chanMap.chMap[topic])
}

//// SendMsgToEs 发送消息到ES服务器；
//func SendMsgToEs(index string, ifLog bool) error {
//	for {
//		msg := kafka.GetMsgFromChan()
//		//msgData, err := json.Marshal(msg)
//		//if err != nil {
//		//	continue
//		//}
//		response, err := cli.Index().Index(index).BodyJson(msg).Do(context.Background())
//		if err != nil {
//			return err
//		}
//		if ifLog {
//			log.Printf("ID %s to index %s, type %s \n", response.Id, response.Index, response.Type)
//		}
//	}
//}
