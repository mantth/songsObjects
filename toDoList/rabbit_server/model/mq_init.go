package model

import (
	"github.com/streadway/amqp"
)

var MQ *amqp.Connection

// RabbitMQ 初始化消息队列；
// notice: RabbitMQ开放了两个端口，之前用错了连不上；
func RabbitMQ(connString string) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	MQ = conn
}
