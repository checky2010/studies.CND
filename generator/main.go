package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Datapoint struct {
	Value int       `json:"value"`
	Date  time.Time `json:"date"`
}

func main() {
	maxValue, err := strconv.Atoi(os.Getenv("MAX_VALUE"))
	if err != nil {
		maxValue = 100
		err = nil
	}
	minValue, err := strconv.Atoi(os.Getenv("MIN_VALUE"))
	if err != nil {
		minValue = 0
		err = nil
	}
	step, err := strconv.Atoi(os.Getenv("STEP"))
	if err != nil {
		step = 5
		err = nil
	}

	channel := getRabbitChannel()

	nextValue := rand.Intn(maxValue)
	for {
		dp := &Datapoint{Value: nextValue, Date: time.Now()}
		jsonDp, err := json.Marshal(dp)
		if err != nil {
			panic("Can't marshal datapoint")
		}
		err = channel.PublishWithContext(
			context.Background(),
			"",
			os.Getenv("RABBITMQ_QUEUE"),
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        jsonDp,
			})
		fmt.Println(dp)

		step := rand.Intn(step*2) - step
		nextValue += step
		if nextValue > maxValue {
			nextValue = maxValue
		}
		if nextValue < minValue {
			nextValue = minValue
		}

		time.Sleep(time.Second)
	}
}

func getRabbitChannel() *amqp.Channel {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		getEnv("RABBITMQ_PORT", "5672"),
	))
	if err != nil {
		println(fmt.Sprintf("amqp://%s:%s@%s:%s/",
			os.Getenv("RABBITMQ_USER"),
			os.Getenv("RABBITMQ_PASSWORD"),
			os.Getenv("RABBITMQ_HOST"),
			getEnv("RABBITMQ_PORT", "5672"),
		))
		panic("Can't connect to RabbitMQ: " + err.Error())
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

	return channel
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
