package rabbitmq

import (
	"encoding/json"
	"evaluation/adapter/rabbitmq/dto"
	"evaluation/domain/model"
	"fmt"
)

type DatapointEventImpl struct {
	Service *Service
}

func (datapointEvents *DatapointEventImpl) ReceiveDatapoints() <-chan *model.Datapoint {
	datapointChannel := make(chan *model.Datapoint)

	go func() {
		for event := range datapointEvents.Service.Receive() {
			fmt.Print("Received event ")
			fmt.Println(event)

			var datapoint dto.NewDatapointEvent
			err := json.Unmarshal(event.Body, &datapoint)
			if err != nil {
				panic("Error unmarshalling datapoint")
			}
			datapointChannel <- &model.Datapoint{
				Value: datapoint.Value,
				Date:  datapoint.Date,
			}
		}
	}()

	return datapointChannel
}
