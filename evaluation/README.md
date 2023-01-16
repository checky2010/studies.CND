# Evaluation

Receives datapoints from a queue and stores them in a database. Also allows for evaluation of the stored data.

## Environment variables

| ENV                 | Default | Optional | Description                                                                    |
|:--------------------|:--------|----------|:-------------------------------------------------------------------------------|
| `PORT`              | _8080_  | Yes      | The port on which the service listens.                                         |
| `RABBIT_URL`        | _-_     | No       | The URL to access RabbitMQ in the format `amqp://username:password@url:5672/`. |
| `RABBIT_QUEUE`      | _-_     | No       | The name of the queue from which the datapoints will be received.              |
| `POSTGRES_HOST`     | _-_     | No       | The host under which Postgres is reachable.                                    |
| `POSTGRES_PORT`     | _-_     | No       | The port under which Postgres is reachable.                                    |
| `POSTGRES_USER`     | _-_     | No       | The user which has access to the database.                                     |
| `POSTGRES_PASSWORD` | _-_     | No       | The password for the user.                                                     |
| `POSTGRES_DATABASE` | _-_     | No       | The database which will be used by the service.                                |

## GraphQL queries

This service supports the following queries:
- `datapoints` -> []Datapoint
- `datapoints(start)` -> []Datapoint
- `datapoints(end)` -> []Datapoint
- `datapoints(start, end)` -> []Datapoint
- `averageValue` -> float
- `averageValue(start)` -> float
- `averageValue(end)` -> float
- `averageValue(start, end)` -> float
- `maxDatapoint` -> Datapoint
- `maxDatapoint(start)` -> Datapoint
- `maxDatapoint(end)` -> Datapoint
- `maxDatapoint(start, end)` -> Datapoint
- `minDatapoint` -> Datapoint
- `minDatapoint(start)` -> Datapoint
- `minDatapoint(end)` -> Datapoint
- `minDatapoint(start, end)` -> Datapoint

A list of test queries which can be pasted into [GraphQL Playground](http://localhost:8080/playground) can be found [here](graphql_queries.txt).

## Endpoints

There are two endpoints this service publishes. First the API at [/api](http://localhost:8080/api) and also a GraphQL Playground at [/playground](http://localhost:8080/playground) which is very useful to send data to the API.