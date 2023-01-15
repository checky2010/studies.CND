package application

import (
	"evaluation/domain/model"
	"evaluation/ports/outgoing"
	"time"
)

type DatapointServiceImpl struct {
	DatapointRepository outgoing.DatapointRepository
}

func (service *DatapointServiceImpl) ReceiveDatapoint(datapointChannel <-chan *model.Datapoint) {
	go func() {
		for datapoint := range datapointChannel {
			service.AddDatapoint(datapoint)
		}
	}()
}

func (service *DatapointServiceImpl) AddDatapoint(datapoint *model.Datapoint) {
	service.DatapointRepository.Save(datapoint)
}

func (service *DatapointServiceImpl) GetDatapointsForTime(start time.Time, end time.Time) []*model.Datapoint {
	return service.DatapointRepository.FindForTime(start, end)
}
