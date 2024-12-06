package rabbitmq

import (
	"github.com/agunfir98/gobroker/lib"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchangeName = "notifications"
)

var queueNames = []string{"email", "sms", "whatsapp"}

type RabbitMq struct {
	connection *amqp.Connection
}

func NewRabbitMq(url string) *RabbitMq {
	conn, err := amqp.Dial(url)
	lib.FailOnError(err, "Failed to connect to Rabbitmq")

	ch, err := conn.Channel()
	lib.FailOnError(err, "failed to connect to the Channel")

	ch.ExchangeDeclare(exchangeName, amqp.ExchangeDirect, true, false, false, false, amqp.Table{})

	for _, name := range queueNames {
		ch.QueueDeclare(name, true, false, false, false, amqp.Table{})
		ch.QueueBind(name, name, exchangeName, false, amqp.Table{})
	}

	rmq := &RabbitMq{
		connection: conn,
	}

	return rmq
}

func (rmq *RabbitMq) Close() {
	rmq.Close()
}
