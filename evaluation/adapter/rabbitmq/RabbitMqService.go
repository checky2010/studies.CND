package rabbitmq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service struct {
	Channel *amqp.Channel
}

func NewServiceImpl() *Service {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		getEnv("RABBITMQ_PORT", "5672"),
	))
	if err != nil {
		panic("Can't connect to RabbitMQ")
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("Can't open channel to RabbitMQ")
	}

	_, err = channel.QueueDeclare(
		os.Getenv("RABBITMQ_QUEUE"),
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		panic("Can't declare queue")
	}

	return &Service{
		Channel: channel,
	}
}

func (service *Service) Receive() <-chan amqp.Delivery {
	events, err := service.Channel.Consume(
		os.Getenv("RABBITMQ_QUEUE"),
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
