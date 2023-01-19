package main

import (
	"evaluation/adapter/api/graphql"
	"evaluation/adapter/gorm"
	"evaluation/adapter/rabbitmq"
	"evaluation/application"
	"evaluation/domain"
	"evaluation/ports/ingoing"
	"evaluation/ports/outgoing"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

func main() {
	var datapointEvents ingoing.DatapointEvent = &rabbitmq.DatapointEventImpl{
		Service: rabbitmq.NewServiceImpl(),
	}

	var datapointRepository outgoing.DatapointRepository = gorm.NewDatapointRepositoryImpl()

	var datapointService domain.DatapointService = &application.DatapointServiceImpl{
		DatapointRepository: datapointRepository,
	}

	datapointService.ReceiveDatapoint(datapointEvents.ReceiveDatapoints())

	mux := http.NewServeMux()
	mux.Handle("/playground", playground.Handler("GraphQL playground", "/api"))
	mux.Handle("/api", handler.NewDefaultServer(
		graphql.NewExecutableSchema(
			graphql.Config{
				Resolvers: &graphql.Resolver{
					StatisticService: &graphql.StatisticServiceImpl{
						DatapointService: datapointService,
					},
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
