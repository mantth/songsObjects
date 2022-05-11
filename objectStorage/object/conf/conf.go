package conf

type RabbitConfig struct {
	RabbitMQHost     string `ini:"RabbitMQHost"`
	RabbitMQPort     string `ini:"RabbitMQPort"`
	RabbitMQPassWord string `ini:"RabbitMQPassWord"`
	RabbitMQUser     string `ini:"RabbitMQUser"`
	RabbitMQ         string `ini:"RabbitMQ"`
}
type Config struct {
	Rabbit *RabbitConfig `ini:"rabbitmq"`
}
