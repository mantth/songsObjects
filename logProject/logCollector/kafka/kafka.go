package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

var (
	kafkaClient sarama.SyncProducer
	// MsgChan 使用指针对象节约内存
	msgChan chan *sarama.ProducerMessage
)

// Init 初始化kafka连接；
func Init(add []string, size int64) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	kafkaClient, err = sarama.NewSyncProducer(add, config)
	if err != nil {
		log.Printf("init kafka conn failed: %s", err)
		return err
	}
	// 初始化消息chan；
	msgChan = make(chan *sarama.ProducerMessage, size)
	// 启用一个协程，将msg发送到kafka;
	go sendMsgToKafka()
	return
}

func sendMsgToKafka() {
	for {
		select {
		case msg := <-msgChan:
			pid, offset, err := kafkaClient.SendMessage(msg)
			if err != nil {
				log.Printf("kafka send msg failed: %s, pid: %v, offset: %v", err, pid, offset)
				return
			}
			log.Println("send to kafka success!", pid, offset)
		}
	}
}

// SendMsgToChan 将消息发送到chan；
// 开放给tails使用，不暴露该channel;
func SendMsgToChan(msg *sarama.ProducerMessage) {
	msgChan <- msg
}
