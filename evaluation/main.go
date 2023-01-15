package main

import (
	"evaluation/adapter/gorm"
	"evaluation/adapter/rabbitmq"
	"evaluation/application"
)

func main() {
	datapointEvents := &rabbitmq.DatapointEventsImpl{
		Service: rabbitmq.NewServiceImpl(),
	}

	datapointService := &application.DatapointServiceImpl{
		DatapointRepository: gorm.NewDatapointRepositoryImpl(),
	}

	datapointService.ReceiveDatapoint(datapointEvents.ReceiveDatapoints())
}
