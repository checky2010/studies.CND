package rabbitmq

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ServiceImpl struct {
	Channel *amqp.Channel
}

func NewServiceImpl() Service {
	connection, err := amqp.Dial(os.Getenv("RABBIT_URL"))
	if err != nil {
		panic("Can't connect to RabbitMQ")
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("Can't open channel to RabbitMQ")
	}

	_, err = channel.QueueDeclare(
		os.Getenv("RABBIT_QUEUE"),
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		panic("Can't declare queue")
	}

	return &ServiceImpl{
		Channel: channel,
	}
}

func (service *ServiceImpl) Receive() <-chan amqp.Delivery {
	events, err := service.Channel.Consume(
		os.Getenv("RABBIT_QUEUE"),
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic("Error consuming from queue")
	}

	return events
}
