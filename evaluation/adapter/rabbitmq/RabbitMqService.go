package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type Service interface {
	Receive() <-chan amqp.Delivery
}
