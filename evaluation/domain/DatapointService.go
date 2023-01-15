package domain

import (
	"evaluation/domain/model"
	"time"
)

type DatapointService interface {
	ReceiveDatapoint(datapointChannel <-chan *model.Datapoint)
	AddDatapoint(datapoint *model.Datapoint)
	GetDatapointsForTime(start, end *time.Time) []*model.Datapoint
}
