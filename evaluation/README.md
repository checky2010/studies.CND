# Evaluation

Receives datapoints from a queue and stores them in a database. Also allows for evaluation of the stored data.

## Environment variables

| ENV                 | Default | Optional | Description                                                                    |
|:--------------------|:--------|----------|:-------------------------------------------------------------------------------|
| `RABBIT_URL`        | _-_     | No       | The URL to access RabbitMQ in the format `amqp://username:password@url:5672/`. |
| `RABBIT_QUEUE`      | _-_     | No       | The name of the queue from which the datapoints will be received.              |
| `POSTGRES_HOST`     | _-_     | No       | The host under which Postgres is reachable.                                    |
| `POSTGRES_PORT`     | _-_     | No       | The port under which Postgres is reachable.                                    |
| `POSTGRES_USER`     | _-_     | No       | The user which has access to the database.                                     |
| `POSTGRES_PASSWORD` | _-_     | No       | The password for the user.                                                     |
| `POSTGRES_DATABASE` | _-_     | No       | The database which will be used by the service.                                |
