package incoming

import "evaluation/domain/model"

type DatapointEvent interface {
	ReceiveDatapoint() *model.Datapoint
}
