# Datapoint visualizer

This project consists of three components: a [generator](generator/README.md), an [evaluator](evaluation/README.md), and a [frontend](frontend/README.md). The generator produces random test data and sends those datapoints to a queue. The evaluator takes them from there and stores them in a database. Now the frontend can request this data and visualize it in four different ways.

## Docker

### RabbitMQ

Start:
```shell
docker run -d \
  -p "5672:5672" \
  -p "8082:15672" \
  --env-file .env \
  -v ./data/dev/rabbitmq:/var/lib/rabbitmq \
  rabbitmq:management
```

### Postgres

Start:
```shell
docker run -d \
  -p "5432:5432" \
  --env-file .env \
  -v ./data/dev/postgres:/var/lib/postgresql/data \
  postgres
```

### Generator

Build:
```shell
cd ./generator
docker build -t generator .
cd ..
```

Start (depends on [rabbitmq](#rabbitmq)):
```shell
docker run -d \
  --env-file .env \
  generator
```

### Evaluation

Build:
```shell
cd ./evaluation
docker build -t evaluation .
cd ..
```

Start (depends on [rabbitmq](#rabbitmq) and [postgres](#postgres)):
```shell
docker run -d \
  -p "8080:8080" \
  --env-file .env \
  evaluation
```

### Frontend

Build:
```shell
cd ./frontend
docker build -t frontend .
cd ..
```

Start (depends on [evaluation](#evaluation)):
```shell
docker run -d \
  -p "8081:80" \
  --env-file .env \
  frontend
```

## Docker Compose

### Development

[Docker Compose File](docker-compose.dev.yml)

This compose file contains all necessary dependencies for the three services, so [RabbitMQ](#rabbitmq) and [Postgresql](#postgres). This file is mainly used for development purposes or for deploying the containers separately without worrying about those dependencies.

### Production

[Docker Compose File](docker-compose.prod.yml)

Start:
```shell
docker compose up
```

After starting the compose file, the frontend is accessible at [http://localhost](http://localhost) and the GraphQL Playground at [http://localhost/playground](http://localhost/playground).

The Docker Compose also contains [traefik](https://traefik.io/traefik/) as a Load-Balancer. All containers can be scaled at will without problems.