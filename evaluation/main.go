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
	"github.com/rs/cors"
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

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/api"))
	mux.Handle("/api", handler.NewDefaultServer(
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

	h := cors.AllowAll().Handler(mux)
	log.Fatal(http.ListenAndServe(":"+port, h))
}
