package service

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"rabbit_server/model"
)

// CreateTask 从消息队列中取出数据在Mysql中创建对应项目；
func CreateTask() {
	ch, err := model.MQ.Channel()
	// 退出时关闭；
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ch)

	if err != nil {
		panic(err)
	}
	queue, err := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	msgs, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// 处于一个监听状态，一致监听我们的生产端的生产；
	go func() {
		for d := range msgs {
			var t model.Task
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			model.DB.Create(&t)
			log.Println("Done")
			_ = d.Ack(false)
		}
	}()

}
