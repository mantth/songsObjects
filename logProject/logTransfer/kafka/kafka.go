package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"logTransfer/es"
)

// Init 初始化kafka连接；
func Init(addr string) (err error) {
	consumer, err := sarama.NewConsumer([]string{addr}, nil)
	// 这里关闭了好像就会出错
	//defer consumer.Close()
	if err != nil {
		return err
	}
	topics, err := consumer.Topics()
	if err != nil {
		log.Println("kafka get topics failed...", err)
		return
	}
	// 为每个topic启用一个协程将消息发送到消息通道；
	for _, topic := range topics {
		partitionList, err := consumer.Partitions(topic)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// kafka与ES异步执行；
		for partition := range partitionList {
			var pc sarama.PartitionConsumer
			pc, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
			if err != nil {
				return err
			}
			go func(sarama.PartitionConsumer) {
				for msg := range pc.Messages() {
					var temp map[string]interface{}
					// 如果出错，则跳过本轮循环；
					if err := json.Unmarshal(msg.Value, &temp); err != nil {
						fmt.Println(err)
						continue
					}
					// 将topic加入消息中；
					temp["topic"] = msg.Topic
					//fmt.Println(temp)
					if es.ChanExist(msg.Topic) {
						es.SendMsgToChan(temp, msg.Topic)
					} else {
						// 前面esMgr中应该可以关闭，因为这里做了判断；
						log.Println("you can't send a msg to a unknown elasticsearch index, please check your etcd config...")
					}
				}
				//defer pc.AsyncClose()
			}(pc)
		}
	}
	return
}

//// GetMsgFromChan 从消息通道中取出消息，开放给ES使用；
//func GetMsgFromChan() interface{} {
//	for {
//		select {
//		case msg := <-msgChan:
//			//fmt.Println(msg.Topic)
//			return msg
//		}
//	}
//
//}
