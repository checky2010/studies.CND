# Generator

Generates random values between 0 and 100 and publishes them to a RabbitMQ queue. The values are always falling/climbing by less than 5 per step.

## Environment variables

| ENV                 | Default | Optional | Description                                                      |
|:--------------------|:--------|----------|:-----------------------------------------------------------------|
| `MAX_VALUE`         | _100_   | Yes      | The maximum value that will be generated.                        |
| `MIN_VALUE`         | _0_     | Yes      | The minimum value that will be generated.                        |
| `STEP`              | _5_     | Yes      | The maximum delta between two data-points                        |
| `RABBITMQ_USERNAME` | _-_     | No       | The username with access to RabbitMQ.                            |
| `RABBITMQ_PASSWORD` | _-_     | No       | The password for the user.                                       |
| `RABBITMQ_HOST`     | _-_     | No       | The URL to access RabbitMQ in the format `localhost`.            |
| `RABBITMQ_PORT`     | _5672_  | Yes      | The port to access RabbitMQ.                                     |
| `RABBITMQ_QUEUE`    | _-_     | No       | The name of the queue where the datapoints will be published to. |
