package domain

import (
	"evaluation/domain/model"
	"time"
)

type DatapointService interface {
	ReceiveDatapoint(datapointChannel <-chan *model.Datapoint)
	AddDatapoint(datapoint *model.Datapoint)
	GetDatapointsForTime(start time.Time, end time.Time) []*model.Datapoint
}
