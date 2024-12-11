package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func (rmq *RabbitMq) Consume(route string) (<-chan amqp.Delivery, error) {
	return rmq.channel.Consume(
		route,
		exchangeName,
		true,
		false,
		false,
		false,
		amqp.Table{},
	)
}
