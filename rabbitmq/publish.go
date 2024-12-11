package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func (rm *RabbitMq) Publish(queueName string, body []byte) error {
	return rm.channel.Publish(
		exchangeName,
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
