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
			os.Getenv("RABBIT_QUEUE"),
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

	return channel
}
