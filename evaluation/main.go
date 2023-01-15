package main

import (
	"evaluation/adapter/api/graphql"
	"evaluation/adapter/gorm"
	"evaluation/adapter/rabbitmq"
	"evaluation/application"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	datapointEvents := &rabbitmq.DatapointEventsImpl{
		Service: rabbitmq.NewServiceImpl(),
	}

	datapointRepository := gorm.NewDatapointRepositoryImpl()

	datapointService := &application.DatapointServiceImpl{
		DatapointRepository: datapointRepository,
	}

	datapointService.ReceiveDatapoint(datapointEvents.ReceiveDatapoints())

	http.Handle("/", playground.Handler("GraphQL playground", "/api"))
	http.Handle("/api", handler.NewDefaultServer(
		graphql.NewExecutableSchema(
			graphql.Config{
				Resolvers: &graphql.Resolver{
					StatisticService: &graphql.StatisticServiceImpl{
						DatapointRepository: datapointRepository,
					},
					DatapointRepository: datapointRepository,
				},
			},
		),
	))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
