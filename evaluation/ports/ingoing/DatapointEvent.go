package ingoing

import "evaluation/domain/model"

type DatapointEvent interface {
	ReceiveDatapoints() <-chan *model.Datapoint
}
