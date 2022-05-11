package model

import "github.com/streadway/amqp"

var MQ *amqp.Connection

func RabbitMQ(connString string) {
	conn, err := amqp.Dial(connString)
	//defer conn.Close()
	if err != nil {
		panic(err)
	}
	MQ = conn
}
