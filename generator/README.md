# Generator

Generates random values between 0 and 100 and publishes them to a RabbitMQ queue.

## Environment variables

| ENV            | Default | Optional | Description                                                                    |
|:---------------|:--------|----------|:-------------------------------------------------------------------------------|
| `MAX_VALUE`    | _100_   | Yes      | The maximum value that will be generated.                                      |
| `MIN_VALUE`    | _0_     | Yes      | The minimum value that will be generated.                                      |
| `STEP`         | _5_     | Yes      | The maximum delta between two data-points                                      |
| `RABBIT_URL`   | _-_     | No       | The URL to access RabbitMQ in the format `amqp://username:password@url:5672/`. |
| `RABBIT_QUEUE` | _-_     | No       | The name of the queue where the datapoints will be published to.               |
